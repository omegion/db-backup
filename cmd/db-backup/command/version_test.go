package command

import (
	"testing"
)

func TestVersionCommand(t *testing.T) {
	_, err := executeCommand(Version())
	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}
