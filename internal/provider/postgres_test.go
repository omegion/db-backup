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
	commander := internal.Commander{}

	p := Postgres{}
	p.SetCommander(commander)

	assert.Equal(t, p.Commander, commander)
}

func TestExport(t *testing.T) {
	post := Postgres{
		Name:     "test_db",
		Host:     "db.example.com",
		Port:     "1234",
		Username: "test_user",
		Password: "123456789",
	}

	bck := backup.Backup{Path: "/var/test/my-bucket-name"}

	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf(
				"pg_dump -d%s -h%s -p%s -U%s -f%s",
				post.Name,
				post.Host,
				post.Port,
				post.Username,
				"my-bucket-name",
			),
		},
	}

	post.Commander = internal.Commander{Executor: test.NewExecutor(expectedCommands)}

	err := post.Export(&bck)

	assert.NoError(t, err)
	assert.Equal(t, post.Password, os.Getenv("PGPASSWORD"))
}

func TestExport_Failure(t *testing.T) {
	post := Postgres{}

	bck := backup.Backup{Path: "/var/test/my-bucket-name"}

	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf(
				"pg_dump -f%s",
				"my-bucket-name",
			),
			StdErr: []byte("custom-error"),
		},
	}

	post.Commander = internal.Commander{Executor: test.NewExecutor(expectedCommands)}

	err := post.Export(&bck)

	assert.EqualError(t, err, "'pg_dump': Execution failed: ")
}

func TestImport(t *testing.T) {
	post := Postgres{
		Name:     "test_db",
		Port:     "1234",
		Username: "test_user",
		Password: "123456789",
	}

	bck := backup.Backup{Path: "/var/test/my-bucket-name"}

	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf(
				"psql -d%s -p%s -U%s -f%s",
				post.Name,
				post.Port,
				post.Username,
				"my-bucket-name",
			),
		},
	}

	post.Commander = internal.Commander{Executor: test.NewExecutor(expectedCommands)}

	err := post.Import(&bck)

	assert.NoError(t, err)
	assert.Equal(t, post.Password, os.Getenv("PGPASSWORD"))
}

func TestImport_Failure(t *testing.T) {
	post := Postgres{}

	bck := backup.Backup{Path: "/var/test/my-bucket-name"}

	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf(
				"psql -f%s",
				"my-bucket-name",
			),
			StdErr: []byte("custom-error"),
		},
	}

	post.Commander = internal.Commander{Executor: test.NewExecutor(expectedCommands)}

	err := post.Import(&bck)

	assert.EqualError(t, err, "'psql': Execution failed: ")
}
