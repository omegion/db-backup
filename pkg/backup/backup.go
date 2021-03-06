package backup

import (
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

// Backup for database.
type Backup struct {
	Name      string
	Path      string
	Host      string
	CreatedAt time.Time
}

func New(options Options) Backup {
	return Backup{
		Name:      options.Name,
		Host:      options.Host,
		CreatedAt: time.Time{},
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
	err := storage.Get(*b)
	if err != nil {
		return err
	}

	return nil
}

// Filename backup filename from path.
func (b *Backup) Filename() string {
	_, filename := filepath.Split(b.Path)

	return filename
}
