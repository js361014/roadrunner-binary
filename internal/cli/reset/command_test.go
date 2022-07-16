package reset_test

import (
	"testing"

	"github.com/js361014/roadrunner-binary/v2/internal/cli/reset"

	"github.com/js361014/roadrunner-plugins/v2/config"
	"github.com/stretchr/testify/assert"
)

func TestCommandProperties(t *testing.T) {
	cmd := reset.NewCommand(&config.Plugin{})

	assert.Equal(t, "reset", cmd.Use)
	assert.NotNil(t, cmd.RunE)
}

func TestExecution(t *testing.T) {
	t.Skip("Command execution is not implemented yet")
}
