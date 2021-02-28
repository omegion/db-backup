package dump

import (
	"github.com/spf13/cobra"
)

func Dump() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dump",
		Short: "dump Management",
		Long:  "CLI command to manage backups",
	}

	cmd.AddCommand(Local())
	cmd.AddCommand(S3())

	return cmd
}
