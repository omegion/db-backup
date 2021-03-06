package database

import (
	"fmt"
	"github.com/omegion/go-command"
	"os"
	"os/exec"
	"time"
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
func (db Postgres) Export() (*Backup, error) {
	backup := &Backup{
		Name: db.Name,
		Host: db.Host,
	}

	t := time.Now()

	backup.Path = fmt.Sprintf(`%v.sql.tar.gz`, t.Format(time.RFC3339))

	options := append(db.dumpOptions(), fmt.Sprintf(`-f%v`, backup.Filename()))

	_, err := db.Commander.Output(PGDumpCmd, options...)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return &Backup{}, PostgresError{
				Origin:  err,
				Message: string(exitError.Stderr),
			}
		}

		return backup, err
	}

	return backup, nil
}

// Import imports given file to Postgres database.
func (db Postgres) Import(file string) (*Backup, error) {
	backup := &Backup{
		Name: db.Name,
		Path: file,
	}

	options := append(db.dumpOptions(), fmt.Sprintf(`-f%v`, backup.Filename()))

	_, err := db.Commander.Output(PGImportCmd, options...)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return &Backup{}, PostgresError{
				Origin:  err,
				Message: string(exitError.Stderr),
			}
		}

		return backup, err
	}

	return backup, nil
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
