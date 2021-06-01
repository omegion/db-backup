package backup

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStorage struct{}

func (s *mockStorage) Get(backup Backup) error         { return nil }
func (s *mockStorage) Save(backup Backup) error        { return nil }
func (s *mockStorage) Delete(backup Backup) error      { return nil }
func (s *mockStorage) List(b Backup) ([]Backup, error) { return []Backup{}, nil }

type mockStorageFailure struct{}

func (s *mockStorageFailure) Get(backup Backup) error    { return errors.New("custom-error") }
func (s *mockStorageFailure) Save(backup Backup) error   { return errors.New("custom-error") }
func (s *mockStorageFailure) Delete(backup Backup) error { return errors.New("custom-error") }
func (s *mockStorageFailure) List(b Backup) ([]Backup, error) {
	return []Backup{}, errors.New("custom-error")
}

func TestBackup_Get(t *testing.T) {
	t.Parallel()

	storage := mockStorage{}
	backup := New(Options{})

	err := backup.Get(&storage)

	assert.NoError(t, err)
}

func TestBackup_Get_Failure(t *testing.T) {
	t.Parallel()

	storage := mockStorageFailure{}
	backup := New(Options{})

	err := backup.Get(&storage)

	assert.EqualError(t, err, "custom-error")
}

func TestBackup_Save(t *testing.T) {
	t.Parallel()

	storage := mockStorage{}
	backup := New(Options{})

	err := backup.Save(&storage)

	assert.NoError(t, err)
}

func TestBackup_Save_Failure(t *testing.T) {
	t.Parallel()

	storage := mockStorageFailure{}
	backup := New(Options{})

	err := backup.Save(&storage)

	assert.EqualError(t, err, "custom-error")
}

func TestBackup_Delete(t *testing.T) {
	t.Parallel()

	storage := mockStorage{}
	backup := New(Options{})

	err := backup.Delete(&storage)

	assert.NoError(t, err)
}

func TestBackup_Delete_Failure(t *testing.T) {
	t.Parallel()

	storage := mockStorageFailure{}
	backup := New(Options{})

	err := backup.Delete(&storage)

	assert.EqualError(t, err, "custom-error")
}

func TestBackup_Filename(t *testing.T) {
	t.Parallel()

	backup := New(Options{})

	fileName := backup.Filename()

	assert.Contains(t, fileName, ".sql.tar.gz")
}
