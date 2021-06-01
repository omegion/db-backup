package provider

import "fmt"

// TypeError occurs when Database type is not supported.
type TypeError struct {
	Type string
}

func (e TypeError) Error() string {
	return fmt.Sprintf("Database type %s is not supported.", e.Type)
}

// ExecutionFailedError occurs when an execution fails.
type ExecutionFailedError struct {
	Command string
	Message string
}

func (e ExecutionFailedError) Error() string {
	return fmt.Sprintf("'%s': Execution failed: %s", e.Command, e.Message)
}
