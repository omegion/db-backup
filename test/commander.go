package test

import (
	"bytes"
	"errors"
	"strings"

	"k8s.io/utils/exec"
	testingexec "k8s.io/utils/exec/testing"
)

// FakeCommand is a command wrapper for testing.
type FakeCommand struct {
	Command string
	StdOut  []byte
	StdErr  []byte
}

// NewExecutor is a factory for Commander testing.
func NewExecutor(commands []FakeCommand) *testingexec.FakeExec {
	cmdActions := make([]testingexec.FakeCommandAction, 0)

	for i := range commands {
		fakeCmd := &commands[i]

		cmdActions = append(cmdActions, func(c string, args ...string) exec.Cmd {
			return &testingexec.FakeCmd{
				Argv: strings.Split(fakeCmd.Command, " "),
				OutputScript: []testingexec.FakeAction{
					func() ([]byte, []byte, error) {
						if bytes.Equal(fakeCmd.StdErr, []byte("")) {
							return fakeCmd.StdOut, nil, nil
						}
						//nolint:goerr113 // allow static errors.
						return fakeCmd.StdOut, nil, errors.New(string(fakeCmd.StdOut))
					},
				},
			}
		})
	}

	return &testingexec.FakeExec{
		ExactOrder:    true,
		CommandScript: cmdActions,
	}
}
