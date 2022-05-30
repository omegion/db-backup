package main

import (
	"os"

	commander "github.com/omegion/cobra-commander"
	"github.com/spf13/cobra"

	"github.com/omegion/db-backup/cmd"
	"github.com/omegion/db-backup/cmd/local"
	"github.com/omegion/db-backup/cmd/s3"
)

func main() {
	root := &cobra.Command{
		Use:          "db-backup",
		Short:        "Database Backup Tool",
		Long:         "CLI command to create backup for provider in local or cloud.",
		SilenceUsage: true,
	}

	comm := commander.NewCommander(root).
		SetCommand(
			cmd.Version(),
			s3.Main(),
			local.Main(),
		).
		Init()

	if err := comm.Execute(); err != nil {
		os.Exit(1)
	}
}
