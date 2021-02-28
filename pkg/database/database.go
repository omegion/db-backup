package database

import "github.com/omegion/go-db-backup/pkg/exporter"

type Database interface {
	Export() (*exporter.ExportResult, error)
}

type Options struct {
	Type     string
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	Options  []string
}
