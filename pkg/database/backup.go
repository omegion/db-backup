package database

import "path/filepath"

// Storage interface of different storages.
type Storage interface {
	Save(backup Backup) error
	Delete(backup Backup) error
	Get(backup Backup) error
	List() // list backup files for given Host and Database
}

// Backup for database.
type Backup struct {
	Name string
	Path string
	Host string
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
