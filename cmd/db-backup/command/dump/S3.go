package dump

import (
	db "github.com/omegion/go-db-backup/pkg/database"
	"github.com/omegion/go-db-backup/pkg/store"
	"github.com/spf13/cobra"
	"log"
)

func setupS3Command(cmd *cobra.Command) {
	cmd.Flags().String("bucket-name", "", "Bucket Name")
	if err := cmd.MarkFlagRequired("bucket-name"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("endpoint-url", "", "Custom Endpoint URL")

}

func S3() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "s3",
		Short: "Dumps database to S3 bucket",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbType, _ := cmd.Flags().GetString("type")
			host, _ := cmd.Flags().GetString("host")
			port, _ := cmd.Flags().GetString("port")
			databaseName, _ := cmd.Flags().GetString("database")
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			bucketName, _ := cmd.Flags().GetString("bucket-name")
			endpointURL, _ := cmd.Flags().GetString("endpoint-url")

			options := db.Options{
				Type:     dbType,
				Host:     host,
				Port:     port,
				Name:     databaseName,
				Username: username,
				Password: password,
				Options:  []string{"--no-owner"},
			}

			database, err := getDatabaseByType(options)
			if err != nil {
				return err
			}

			result, err := database.Export()
			if err != nil {
				return err
			}

			export := store.S3{
				Bucket:      bucketName,
				EndpointURL: endpointURL,
			}

			err = result.To(&export)
			if err != nil {
				return err
			}

			return nil
		},
	}

	setupDumpCommand(cmd)
	setupS3Command(cmd)

	return cmd
}
