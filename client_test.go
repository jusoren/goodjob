package goodjob_test

import (
	"testing"

	"github.com/jusoren/goodjob"
	"github.com/stretchr/testify/assert"
)

func TestCreateJobEmpty(t *testing.T) {
	var err error

	job := goodjob.Job{}
	err = client.CreateJob(&job)

	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, err.Error(), "name is required")
	}
}

func TestCreateJobWithJustName(t *testing.T) {
	var err error

	job := goodjob.Job{
		Name: "test-client-1",
	}
	err = client.CreateJob(&job)

	assert.Nil(t, err)
	assert.NotEmpty(t, job.ID)
	assert.Equal(t, "test-client-1", job.Name)
	assert.Equal(t, 60, job.Timeout)
	assert.Equal(t, 3, job.MaxRetries)
}

func TestCreateJobWithNameAndData(t *testing.T) {
	var err error

	job := goodjob.Job{
		Name: "test-client-1",
		Data: `{"foo": "bar"}`,
	}
	err = client.CreateJob(&job)

	assert.Nil(t, err)
	assert.NotEmpty(t, job.ID)
	assert.Equal(t, "test-client-1", job.Name)
	assert.Equal(t, `{"foo": "bar"}`, job.Data)
}

func TestCreateJobWithNoRetry(t *testing.T) {
	var err error

	job := goodjob.Job{
		Name:       "test-client-2",
		MaxRetries: goodjob.NoRetry,
	}
	err = client.CreateJob(&job)

	assert.Nil(t, err)
	assert.Equal(t, goodjob.NoRetry, job.MaxRetries)
}

func TestCreateJobWithMaxRetries7(t *testing.T) {
	var err error

	job := goodjob.Job{
		Name:       "test-client-2",
		MaxRetries: 7,
	}
	err = client.CreateJob(&job)

	assert.Nil(t, err)
	assert.Equal(t, 7, job.MaxRetries)
}
