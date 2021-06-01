package backup

import (
	"fmt"
	"path/filepath"
	"time"
)

// Storage interface of different storages.
type Storage interface {
	Save(Backup) error
	Delete(Backup) error
	Get(Backup) error
	List(Backup) ([]Backup, error)
}

// Options options for Backup.
type Options struct {
	Name string
	Path string
	Host string
}

// Backup for provider.
type Backup struct {
	Name      string
	Path      string
	Host      string
	CreatedAt time.Time
}

// New returns backup filled with Options.
func New(options Options) Backup {
	now := time.Now()

	return Backup{
		Name:      options.Name,
		Host:      options.Host,
		CreatedAt: now,
		Path:      fmt.Sprintf("%v.sql.tar.gz", now.Format(time.RFC3339)),
	}
}

// Save backup to given storage.
func (b *Backup) Save(storage Storage) error {
	return storage.Save(*b)
}

// Delete backup from given storage.
func (b *Backup) Delete(storage Storage) error {
	return storage.Delete(*b)
}

// Get backup from given storage.
func (b *Backup) Get(storage Storage) error {
	return storage.Get(*b)
}

// Filename backup filename from path.
func (b *Backup) Filename() string {
	_, filename := filepath.Split(b.Path)

	return filename
}
