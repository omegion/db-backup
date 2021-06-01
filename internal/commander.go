package internal

import (
	"bytes"
	"errors"
	"k8s.io/utils/exec"
	testingexec "k8s.io/utils/exec/testing"
	"strings"
)

// Commander is the root of this helper.
type Commander struct {
	Executor exec.Interface
}

// NewCommander is a factory for Commander.
func NewCommander() Commander {
	return Commander{
		exec.New(),
	}
}

type FakeCommand struct {
	Command string
	StdOut  []byte
	StdErr  []byte
}

func NewExecutor(commands []FakeCommand) *testingexec.FakeExec {
	cmdActions := make([]testingexec.FakeCommandAction, 0)

	for _, command := range commands {
		cmdActions = append(cmdActions, func(cmd string, args ...string) exec.Cmd {
			holder := &testingexec.FakeCmd{
				Argv: strings.Split(command.Command, " "),
				OutputScript: []testingexec.FakeAction{
					func() ([]byte, []byte, error) {
						if bytes.Compare(command.StdErr, []byte("")) == 0 {
							return command.StdOut, nil, nil
						}
						return command.StdOut, command.StdErr, errors.New(string(command.StdOut))
					},
				},
			}
			buf := bytes.NewBuffer(command.StdErr)
			holder.SetStderr(buf)
			return holder
		})
	}

	return &testingexec.FakeExec{
		ExactOrder:    true,
		CommandScript: cmdActions,
	}
}
