package local

import (
	"fmt"
	"log"
	"strings"

	"github.com/omegion/db-backup/internal"
	"github.com/omegion/db-backup/internal/backup"

	"github.com/spf13/cobra"
)

func setupImportCommand(cmd *cobra.Command) {
	cmd.Flags().String("file", "", "Backup file path")

	if err := cmd.MarkFlagRequired("file"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}
}

// Import imports given backups to provider.
func Import() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import provider backup from local",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbType, _ := cmd.Flags().GetString("type")
			file, _ := cmd.Flags().GetString("file")
			host, _ := cmd.Flags().GetString("host")
			port, _ := cmd.Flags().GetString("port")
			databases, _ := cmd.Flags().GetString("databases")
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")

			commander := internal.NewCommander()

			for _, databaseName := range strings.Split(databases, ",") {
				options := internal.Options{
					Type:     dbType,
					Host:     host,
					Port:     port,
					Name:     databaseName,
					Username: username,
					Password: password,
				}

				database, err := GetDatabaseByType(options)
				if err != nil {
					return err
				}

				database.SetCommander(commander)

				backupOptions := backup.Options{
					Name: databaseName,
					Host: host,
					Path: file,
				}

				b := backup.New(backupOptions)

				err = database.Import(&b)
				if err != nil {
					return err
				}

				fmt.Printf("Database %s imported successfully.\n", databaseName)
			}

			return nil
		},
	}

	setupImportCommand(cmd)
	SetupExportCommand(cmd)

	return cmd
}
