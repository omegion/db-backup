package command

import (
	"github.com/omegion/go-db-backup/cmd/db-backup/command/local"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	_, err := executeCommand(local.Import(),
		"--type=postgres",
		"--file=2021-03-01T20:57:05+01:00.sql.tar.gz",
		"--database=vault1",
		"--host=db.omegion.dev",
		"--password=p6tNYH*4UmrWCP&C9rB#5RdM",
		"--port=2052",
		"--username=omegion",
	)

	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}
