package s3

import (
	"log"

	"github.com/omegion/go-db-backup/cmd/db-backup/command/local"
	db "github.com/omegion/go-db-backup/pkg/database"
	"github.com/omegion/go-db-backup/pkg/storage"
	"github.com/spf13/cobra"
)

func setupExportCommand(cmd *cobra.Command) {
	cmd.Flags().String("bucket", "", "Bucket name")
	if err := cmd.MarkFlagRequired("bucket"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("endpoint", "", "S3 custom endpoint")
	if err := cmd.MarkFlagRequired("endpoint"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}
}

func Export() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export database to S3 bucket.",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbType, _ := cmd.Flags().GetString("type")
			host, _ := cmd.Flags().GetString("host")
			port, _ := cmd.Flags().GetString("port")
			databaseName, _ := cmd.Flags().GetString("database")
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			bucketName, _ := cmd.Flags().GetString("bucket")
			endpointURL, _ := cmd.Flags().GetString("endpoint")

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

			backup, err := database.Export()
			if err != nil {
				return err
			}

			err = backup.Save(&storage.S3{
				Bucket:      bucketName,
				EndpointURL: endpointURL,
			})
			if err != nil {
				return err
			}

			return nil
		},
	}

	local.SetupExportCommand(cmd)
	setupExportCommand(cmd)

	return cmd
}
