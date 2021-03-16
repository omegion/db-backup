package database

import (
	"github.com/omegion/db-backup/pkg/backup"

	"github.com/omegion/go-command"
)

// Database interface for different databases.
type Database interface {
	SetCommander(command.Interface)
	Export(*backup.Backup) error
	Import(*backup.Backup) error
}

// Options for different Database.
type Options struct {
	Type     string
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	Options  []string
}
