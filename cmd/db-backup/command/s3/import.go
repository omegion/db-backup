package s3

import (
	"fmt"
	"github.com/omegion/go-db-backup/cmd/db-backup/command/local"
	db "github.com/omegion/go-db-backup/pkg/database"
	"github.com/omegion/go-db-backup/pkg/storage"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func setupImportCommand(cmd *cobra.Command) {
	cmd.Flags().String("path", "", "Backup path in S3")
	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}
}

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

				backup := db.Backup{
					Name: databaseName,
					Path: path,
					Host: host,
				}

				err = backup.Get(&storage.S3{
					Bucket:      bucketName,
					EndpointURL: endpointURL,
				})
				if err != nil {
					return err
				}

				_, err = database.Import(backup.Path)
				if err != nil {
					return err
				}

				fmt.Print(fmt.Sprintf("Database %s imported successfully.\n", databaseName))
			}

			return nil
		},
	}

	local.SetupExportCommand(cmd)
	setupExportCommand(cmd)
	setupImportCommand(cmd)

	return cmd
}
