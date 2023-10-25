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
		ID:        cuid.New(),
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Timeout:   60,
	}

	if err := mergo.Merge(&job, defaultJob, mergo.WithOverride); err != nil {
		return err
	}

	return c.Driver.CreateJob(job)
}
