package internal

import (
	"github.com/omegion/db-backup/internal/backup"
)

// Database interface for different databases.
type Database interface {
	SetCommander(commander Commander)
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
