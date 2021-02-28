package store

import (
	"fmt"
	"github.com/omegion/go-db-backup/pkg/exporter"
	"os/exec"
	"path/filepath"
)

type Local struct {
	Directory string
}

func (l *Local) Store(backup *exporter.ExportResult) error {
	path := filepath.Join(l.Directory, backup.Filename())
	_, err := exec.Command("mv", backup.Path, path).Output()

	fmt.Println(fmt.Sprintf("Dump is successful for %s", backup.DatabaseName))

	return err
}
