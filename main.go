package main

import (
	cmd2 "github.com/omegion/db-backup/cmd"
	"github.com/omegion/db-backup/cmd/local"
	"github.com/omegion/db-backup/cmd/s3"

	"github.com/spf13/cobra"
)

// RootCommand is the main entry point of this application.
func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db-backup",
		Short: "Kubernetes as a Service Cluster Command",
		Long:  "CLI command to generate configs for Kubernetes as a Cluster",
	}

	cmd.AddCommand(cmd2.Version())
	cmd.AddCommand(s3.Main())
	cmd.AddCommand(local.Main())

	return cmd
}

func main() {
	root := RootCommand()
	_ = root.Execute()
}
