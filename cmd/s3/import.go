package s3

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/db-backup/cmd/local"
	"github.com/omegion/db-backup/internal"
	"github.com/omegion/db-backup/internal/backup"
	"github.com/omegion/db-backup/internal/storage"
)

func setupImportCommand(cmd *cobra.Command) {
	cmd.Flags().String("path", "", "Backup path in S3")

	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}
}

// Import imports given backups to provider.
func Import() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import provider backup from S3 bucket.",
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

				database, err := local.GetDatabaseByType(options)
				if err != nil {
					return err
				}

				database.SetCommander(commander)

				bck := backup.Backup{
					Name: databaseName,
					Path: path,
					Host: host,
				}

				err = bck.Get(&storage.S3{
					Bucket:      bucketName,
					EndpointURL: endpointURL,
				})
				if err != nil {
					return err
				}

				err = database.Import(&bck)
				if err != nil {
					return err
				}

				log.Infoln(fmt.Sprintf("Database %s imported successfully.", databaseName))
			}

			return nil
		},
	}

	local.SetupExportCommand(cmd)
	setupExportCommand(cmd)
	setupImportCommand(cmd)

	return cmd
}
