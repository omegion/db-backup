package database

type Database interface {
	Export() (*Backup, error)
	Import(file string) (*Backup, error)
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
