package database

import "fmt"

// TypeError occurs when Database type is not supported.
type TypeError struct {
	Type string
}

func (e TypeError) Error() string {
	return fmt.Sprintf("Database type %s is not supported.", e.Type)
}

// PostgresError occurs when Postgres command fails.
type PostgresError struct {
	Origin  error
	Message string
}

func (e PostgresError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Origin.Error())
}
