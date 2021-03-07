package database

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/omegion/go-db-backup/pkg/backup"

	"github.com/omegion/go-command"
)

const (
	// PGDumpCmd default dump command for Postgres.
	PGDumpCmd = "pg_dump"
	// PGImportCmd default import command for Postgres.
	PGImportCmd = "psql"
)

// Postgres database.
type Postgres struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	// Extra pg_dump options
	// e.g []string{"--inserts"}
	Options   []string
	Commander command.Interface
}

// SetCommander sets commander for Postgres.
func (db *Postgres) SetCommander(commander command.Interface) {
	db.Commander = commander
}

// Export returns backup with dumped file.
func (db Postgres) Export(backup *backup.Backup) error {
	options := append(db.dumpOptions(), fmt.Sprintf(`-f%v`, backup.Filename()))

	_, err := db.Commander.Output(PGDumpCmd, options...)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return PostgresError{
				Origin:  err,
				Message: string(exitError.Stderr),
			}
		}

		return err
	}

	return nil
}

// Import imports given file to Postgres database.
func (db Postgres) Import(backup *backup.Backup) error {
	options := append(db.dumpOptions(), fmt.Sprintf(`-f%v`, backup.Filename()))

	_, err := db.Commander.Output(PGImportCmd, options...)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return PostgresError{
				Origin:  err,
				Message: string(exitError.Stderr),
			}
		}

		return err
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
