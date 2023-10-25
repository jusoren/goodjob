package goodjob

import (
	"time"
)

const NoRetry = -1

type Tabler interface {
	TableName() string
}

type Job struct {
	ID         string `gorm:"primaryKey"`
	Name       string
	Status     string
	Data       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Timeout    int
	MaxRetries int
}

func (Job) TableName() string {
	return "goodjob_jobs"
}

type Execution struct {
	ID        string `gorm:"primaryKey"`
	Status    string
	Result    *string
	StartedAt time.Time
	EndedAt   *time.Time
	TimeoutAt time.Time
	Retry     uint
	JobID     string
	Job       Job `gorm:"foreignKey:JobID"`
}

func (Execution) TableName() string {
	return "goodjob_executions"
}
