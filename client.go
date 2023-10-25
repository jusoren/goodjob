package goodjob

import (
	"time"

	"dario.cat/mergo"
	"github.com/lucsky/cuid"
)

type Client struct {
	Driver Driver
}

func NewClient(driver Driver) *Client {
	return &Client{
		Driver: driver,
	}
}

func (c *Client) CreateJob(job *Job) error {
	defaultJob := Job{
		ID:         cuid.New(),
		Status:     "pending",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Timeout:    60,
		MaxRetries: 3,
	}

	maxRetries := job.MaxRetries

	if err := mergo.Merge(job, defaultJob, mergo.WithOverride); err != nil {
		return err
	}

	if maxRetries != 0 {
		job.MaxRetries = maxRetries
	}

	return c.Driver.CreateJob(job)
}
