package provider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/omegion/db-backup/internal"
	"github.com/omegion/db-backup/internal/backup"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	// PGDumpCmd default dump command for Postgres.
	PGDumpCmd = "pg_dump"
	// PGImportCmd default import command for Postgres.
	PGImportCmd = "psql"
)

// Postgres provider.
type Postgres struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	// Extra pg_dump options
	// e.g []string{"--inserts"}
	Options   []string
	Commander internal.Commander
}

// SetCommander sets commander for Postgres.
func (db *Postgres) SetCommander(commander internal.Commander) {
	db.Commander = commander
}

// Export returns backup with dumped file.
func (db Postgres) Export(backup *backup.Backup) error {
	options := append(db.dumpOptions(), fmt.Sprintf(`-f%v`, backup.Filename()))

	log.Debugln(fmt.Sprintf("Exporting database table with %s", PGDumpCmd))
	command := db.Commander.Executor.CommandContext(
		context.Background(),
		PGDumpCmd,
		options...,
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	_, err := command.Output()
	if err != nil {
		return ExecutionFailedError{Command: PGDumpCmd, Message: stderr.String()}
	}

	return nil
}

// Import imports given file to Postgres provider.
func (db Postgres) Import(backup *backup.Backup) error {
	options := append(db.dumpOptions(), fmt.Sprintf(`-f%v`, backup.Filename()))

	command := db.Commander.Executor.CommandContext(
		context.Background(),
		PGImportCmd,
		options...,
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	_, err := command.Output()
	if err != nil {
		return ExecutionFailedError{Command: PGImportCmd, Message: stderr.String()}
	}

	return nil
}

func (db Postgres) dumpOptions() []string {
	options := db.Options

	if db.Name != "" {
		options = append(options, fmt.Sprintf(`-d%v`, db.Name))
	}

	if db.Host != "" {
		options = append(options, fmt.Sprintf(`-h%v`, db.Host))
	}

	if db.Port != "" {
		options = append(options, fmt.Sprintf(`-p%v`, db.Port))
	}

	if db.Username != "" {
		options = append(options, fmt.Sprintf(`-U%v`, db.Username))
	}

	if db.Password != "" {
		_ = os.Setenv("PGPASSWORD", db.Password)
	}

	return options
}
