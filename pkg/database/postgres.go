package database

import (
	"fmt"
	"github.com/omegion/go-db-backup/pkg/exporter"
	"os"
	"os/exec"
	"time"
)

var (
	PGDumpCmd = "pg_dump"
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

// Export produces a `pg_dump` of the specified database, and creates a gzip compressed tarball archive.
func (db Postgres) Export() (*exporter.ExportResult, error) {
	result := &exporter.ExportResult{
		MIME:         "application/db-tar",
		DatabaseName: db.Name,
	}

	t := time.Now()

	result.Path = fmt.Sprintf(`%v_%v.sql.tar.gz`, db.Name, t.Format(time.RFC3339))

	options := append(db.dumpOptions(), "-Fc", fmt.Sprintf(`-f%v`, result.Path))

	out, err := exec.Command(PGDumpCmd, options...).Output()
	if err != nil {
		return &exporter.ExportResult{}, PostgresError{
			Origin:  err,
			Message: string(out),
		}
	}

	return result, nil
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
