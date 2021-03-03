package main

import (
	"github.com/omegion/go-db-backup/cmd/db-backup/command"
	"github.com/omegion/go-db-backup/cmd/db-backup/command/local"
	"github.com/omegion/go-db-backup/cmd/db-backup/command/s3"

	"github.com/spf13/cobra"
)

// RootCommand is the main entry point of this application.
func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db-backup",
		Short: "Kubernetes as a Service Cluster Command",
		Long:  "CLI command to generate configs for Kubernetes as a Cluster",
	}

	cmd.AddCommand(command.Version())
	cmd.AddCommand(s3.Main())
	cmd.AddCommand(local.Main())

	return cmd
}

func main() {
	root := RootCommand()
	_ = root.Execute()
}
