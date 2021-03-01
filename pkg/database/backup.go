package database

import "path/filepath"

type Storage interface {
	Save(backup Backup) error
	Delete(backup Backup) error
	Get(backup Backup) error
	List() // list backup files for given Host and Database
}

type Backup struct {
	Name string
	Path string
	Host string
}

func (b *Backup) Save(storage Storage) error {
	return storage.Save(*b)
}

func (b *Backup) Delete(storage Storage) error {
	return storage.Delete(*b)
}

func (b *Backup) Get(storage Storage) error {
	err := storage.Get(*b)
	if err != nil {
		return err
	}

	return nil
}

func (b *Backup) Filename() string {
	_, filename := filepath.Split(b.Path)

	return filename
}
