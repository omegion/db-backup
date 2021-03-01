package database

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	PGDumpCmd   = "pg_dump"
	PGImportCmd = "psql"
)

type Postgres struct {
	// Database Host (e.g. 127.0.0.1)
	Host string
	// Database Port (e.g. 5432)
	Port string
	// Database Name
	Name string
	// Connection Username
	Username string
	// Connection Password
	Password string
	// Extra pg_dump options
	// e.g []string{"--inserts"}
	Options []string
}

func (db Postgres) Export() (*Backup, error) {
	backup := &Backup{
		Name: db.Name,
		Host: db.Host,
	}

	t := time.Now()

	backup.Path = fmt.Sprintf(`%v.sql.tar.gz`, t.Format(time.RFC3339))

	options := append(db.dumpOptions(), fmt.Sprintf(`-f%v`, backup.Filename()))

	out, err := exec.Command(PGDumpCmd, options...).Output()
	if err != nil {
		return &Backup{}, PostgresError{
			Origin:  err,
			Message: string(out),
		}
	}

	return backup, nil
}

func (db Postgres) Import(file string) (*Backup, error) {
	backup := &Backup{
		Name: db.Name,
		Path: file,
	}

	options := append(db.dumpOptions(), fmt.Sprintf(`-f%v`, backup.Filename()))

	_, err := exec.Command(PGImportCmd, options...).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return &Backup{}, PostgresError{
				Origin:  err,
				Message: string(exitError.Stderr),
			}
		}
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
