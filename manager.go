package goodjob

import (
	"errors"
	"time"

	"dario.cat/mergo"
	"github.com/lucsky/cuid"
)

type Manager struct {
	Driver    Driver
	Executors map[string]Executor
}

type Executor struct {
	JobName       string
	Func          func(data string)
	BatchSize     int
	MaxRetries    int
	MaxConcurrent int
}

func NewManager(driver Driver) *Manager {
	return &Manager{
		Driver:    driver,
		Executors: make(map[string]Executor),
	}
}

func (m *Manager) SetExecutor(e Executor) error {
	defaultExecutor := Executor{
		BatchSize: 10,
	}

	if err := mergo.Merge(&e, defaultExecutor, mergo.WithOverride); err != nil {
		return err
	}

	m.Executors[e.JobName] = e

	return nil
}

func (m *Manager) GetExecutor(jobName string) (*Executor, error) {
	executor, ok := m.Executors[jobName]
	if !ok {
		return nil, errors.New("executor not found")
	}

	return &executor, nil
}

func (m *Manager) Run() {
	for {
		jobs, err := m.Driver.GetPendingTimeoutAndErrorJobs()
		if err != nil {
			panic(err)
		}

		for _, job := range jobs {
			executor, ok := m.Executors[job.Name]
			if !ok {
				panic("executor not found")
			}

			execution := Execution{
				ID:        cuid.New(),
				Status:    "running",
				StartedAt: time.Now(),
				TimeoutAt: time.Now().Add(time.Duration(job.Timeout) * time.Second),
				JobID:     job.ID,
			}

			err = m.Driver.CreateExecution(&execution)
			if err != nil {
				panic(err)
			}

			go func() {
				defer func() {
					if r := recover(); r != nil {
						execution.Status = "failed"
						execution.EndedAt = &time.Time{}
						execution.Result = nil
						execution.Retry = execution.Retry + 1
						m.Driver.UpdateExecution(&execution)
					}
				}()

				executor.Func(job.Data)

				execution.Status = "success"
				execution.EndedAt = &time.Time{}
				execution.Result = nil
				execution.Retry = execution.Retry + 1
				m.Driver.UpdateExecution(&execution)
			}()
		}

		time.Sleep(1 * time.Second)
	}
}
