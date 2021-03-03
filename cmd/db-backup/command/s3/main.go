package s3

import "github.com/spf13/cobra"

// Main holds s3 subcommands.
func Main() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "s3",
		Short: "dump Management",
		Long:  "CLI command to manage backups",
	}

	cmd.AddCommand(Export())
	cmd.AddCommand(Import())

	return cmd
}
