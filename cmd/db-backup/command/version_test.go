package command

import (
	"testing"
)

func TestTerraformConfig(t *testing.T) {
	_, err := executeCommand(Version())

	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}
