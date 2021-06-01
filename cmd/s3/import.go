package s3

import (
	"fmt"
	"log"
	"strings"

	"github.com/omegion/db-backup/cmd/local"
	"github.com/omegion/db-backup/internal/backup"
	db "github.com/omegion/db-backup/internal/database"
	"github.com/omegion/db-backup/internal/storage"

	"github.com/omegion/go-command"
	"github.com/spf13/cobra"
)

func setupImportCommand(cmd *cobra.Command) {
	cmd.Flags().String("path", "", "Backup path in S3")

	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}
}

// Import imports given backups to database.
func Import() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import database backup from S3 bucket.",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbType, _ := cmd.Flags().GetString("type")
			path, _ := cmd.Flags().GetString("path")
			host, _ := cmd.Flags().GetString("host")
			port, _ := cmd.Flags().GetString("port")
			databases, _ := cmd.Flags().GetString("databases")
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			bucketName, _ := cmd.Flags().GetString("bucket")
			endpointURL, _ := cmd.Flags().GetString("endpoint")

			commander := command.Command{}

			for _, databaseName := range strings.Split(databases, ",") {
				options := db.Options{
					Type:     dbType,
					Host:     host,
					Port:     port,
					Name:     databaseName,
					Username: username,
					Password: password,
				}

				database, err := local.GetDatabaseByType(options)
				if err != nil {
					return err
				}

				database.SetCommander(commander)

				b := backup.Backup{
					Name: databaseName,
					Path: path,
					Host: host,
				}

				err = b.Get(&storage.S3{
					Bucket:      bucketName,
					EndpointURL: endpointURL,
				})
				if err != nil {
					return err
				}

				err = database.Import(&b)
				if err != nil {
					return err
				}

				fmt.Printf("Database %s imported successfully.\n", databaseName)
			}

			return nil
		},
	}

	local.SetupExportCommand(cmd)
	setupExportCommand(cmd)
	setupImportCommand(cmd)

	return cmd
}
