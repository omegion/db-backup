package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/omegion/db-backup/internal"
	"github.com/omegion/db-backup/internal/backup"
	"github.com/omegion/db-backup/test"
)

func TestSetCommander(t *testing.T) {
	t.Parallel()

	commander := internal.Commander{}

	p := Postgres{}
	p.SetCommander(commander)

	assert.Equal(t, p.Commander, commander)
}

func TestExport(t *testing.T) {
	t.Parallel()

	p := Postgres{
		Name:     "test_db",
		Host:     "db.example.com",
		Port:     "1234",
		Username: "test_user",
		Password: "123456789",
	}

	b := backup.Backup{Path: "/var/test/my-bucket-name"}

	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf(
				"pg_dump -d%s -h%s -p%s -U%s -f%s",
				p.Name,
				p.Host,
				p.Port,
				p.Username,
				"my-bucket-name",
			),
		},
	}

	p.Commander = internal.Commander{Executor: test.NewExecutor(expectedCommands)}

	err := p.Export(&b)

	assert.NoError(t, err)
	assert.Equal(t, p.Password, os.Getenv("PGPASSWORD"))
}

func TestExport_Failure(t *testing.T) {
	t.Parallel()

	p := Postgres{}

	b := backup.Backup{Path: "/var/test/my-bucket-name"}

	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf(
				"pg_dump -f%s",
				"my-bucket-name",
			),
			StdErr: []byte("custom-error"),
		},
	}

	p.Commander = internal.Commander{Executor: test.NewExecutor(expectedCommands)}

	err := p.Export(&b)

	assert.EqualError(t, err, "'pg_dump': Execution failed: ")
}

func TestImport(t *testing.T) {
	t.Parallel()

	p := Postgres{
		Name:     "test_db",
		Port:     "1234",
		Username: "test_user",
		Password: "123456789",
	}

	b := backup.Backup{Path: "/var/test/my-bucket-name"}

	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf(
				"psql -d%s -p%s -U%s -f%s",
				p.Name,
				p.Port,
				p.Username,
				"my-bucket-name",
			),
		},
	}

	p.Commander = internal.Commander{Executor: test.NewExecutor(expectedCommands)}

	err := p.Import(&b)

	assert.NoError(t, err)
	assert.Equal(t, p.Password, os.Getenv("PGPASSWORD"))
}

func TestImport_Failure(t *testing.T) {
	t.Parallel()

	p := Postgres{}

	b := backup.Backup{Path: "/var/test/my-bucket-name"}

	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf(
				"psql -f%s",
				"my-bucket-name",
			),
			StdErr: []byte("custom-error"),
		},
	}

	p.Commander = internal.Commander{Executor: test.NewExecutor(expectedCommands)}

	err := p.Import(&b)

	assert.EqualError(t, err, "'psql': Execution failed: ")
}
