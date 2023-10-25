package goodjob

import "time"

type Manager struct {
	Driver         *Driver
	Executors      map[string]Executor
	DefaultTimeout time.Duration
}

type Executor struct {
	JobName string
	Func    func(data string)
}

func NewManager(driver *Driver) *Manager {
	return &Manager{
		Driver:    driver,
		Executors: make(map[string]Executor),
	}
}

func (m *Manager) SetExecutor(e Executor) {
	m.Executors[e.JobName] = e
}

func (m *Manager) Run() {}
