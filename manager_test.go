package goodjob_test

import (
	"testing"

	"github.com/jusoren/goodjob"
	"github.com/stretchr/testify/assert"
)

func TestSetExecutor(t *testing.T) {
	err := manager.SetExecutor(goodjob.Executor{
		JobName: "test-manager-1",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, manager.Executors["test-manager-1"])
}
