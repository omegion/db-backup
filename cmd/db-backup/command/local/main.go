package local

import "github.com/spf13/cobra"

// Main holds subcommands for local.
func Main() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "local",
		Short: "dump Management",
		Long:  "CLI command to manage backups",
	}

	cmd.AddCommand(Export())
	cmd.AddCommand(Import())

	return cmd
}
