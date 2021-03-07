package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	output, err := executeCommand(Version())
	if err != nil {
		t.Errorf("Command Error: %v", err)
	}

	assert.Equal(t, output, "")
}
