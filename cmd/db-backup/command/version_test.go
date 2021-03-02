package command

import (
	"github.com/omegion/go-db-backup/cmd/db-backup/command/local"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	_, err := executeCommand(local.Export(),
		"--type=postgres",
		"--databases=test",
		"--host=example.com",
		"--password=X",
		"--port=1234",
		"--username=test",
	)

	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}
