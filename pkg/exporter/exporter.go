package exporter

import (
	"os"
	"path/filepath"
)

type Store interface {
	Store(result *ExportResult) error
}

type Exporter interface {
	Export() (*ExportResult, error)
}

type ExportResult struct {
	Path         string
	MIME         string
	DatabaseName string
}

func (e *ExportResult) To(store Store) error {
	err := store.Store(e)
	if err != nil {
		return err
	}

	if _, err := os.Stat(e.Path); os.IsExist(err) {
		err = os.Remove(e.Path)

		return err
	}

	return err
}

func (e ExportResult) Filename() string {
	_, filename := filepath.Split(e.Path)
	return filename
}
