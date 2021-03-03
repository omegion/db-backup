package database

// Database interface for different databases.
type Database interface {
	Export() (*Backup, error)
	Import(file string) (*Backup, error)
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
