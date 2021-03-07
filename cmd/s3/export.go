package s3

import (
	"fmt"
	"log"
	"strings"

	"github.com/omegion/go-db-backup/cmd/local"
	"github.com/omegion/go-db-backup/pkg/backup"
	db "github.com/omegion/go-db-backup/pkg/database"
	"github.com/omegion/go-db-backup/pkg/storage"

	"github.com/omegion/go-command"
	"github.com/spf13/cobra"
)

func setupExportCommand(cmd *cobra.Command) {
	cmd.Flags().String("bucket", "", "Bucket name")

	if err := cmd.MarkFlagRequired("bucket"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("endpoint", "", "S3 custom endpoint")
}

// Export exports given tables from database.
func Export() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export database to S3 bucket.",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbType, _ := cmd.Flags().GetString("type")
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

				backupOptions := backup.Options{
					Name: databaseName,
					Host: host,
				}

				b := backup.New(backupOptions)

				err = database.Export(&b)
				if err != nil {
					return err
				}

				err = b.Save(&storage.S3{
					Bucket:      bucketName,
					EndpointURL: endpointURL,
				})
				if err != nil {
					return err
				}

				fmt.Printf("Database %s exported successfully.\n", databaseName)
			}

			return nil
		},
	}

	local.SetupExportCommand(cmd)
	setupExportCommand(cmd)

	return cmd
}
