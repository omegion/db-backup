package local

import (
	"fmt"
	"log"
	"strings"

	"github.com/omegion/db-backup/internal"
	"github.com/omegion/db-backup/internal/backup"
	db "github.com/omegion/db-backup/internal/provider"

	"github.com/spf13/cobra"
)

// SetupExportCommand sets default flags.
func SetupExportCommand(cmd *cobra.Command) {
	cmd.Flags().String("type", "postgres", "Database type")

	cmd.Flags().String("host", "", "Host")

	if err := cmd.MarkFlagRequired("host"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("port", "", "Port")

	if err := cmd.MarkFlagRequired("port"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("databases", "", "Databases name, e.g. foo,boo")

	if err := cmd.MarkFlagRequired("databases"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("username", "", "Username")

	if err := cmd.MarkFlagRequired("username"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("password", "", "Password")

	if err := cmd.MarkFlagRequired("password"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}
}

// GetDatabaseByType gets provider by its type.
func GetDatabaseByType(options internal.Options) (internal.Database, error) {
	if options.Type == "postgres" {
		return &db.Postgres{
			Host:     options.Host,
			Port:     options.Port,
			Name:     options.Name,
			Username: options.Username,
			Password: options.Password,
			Options:  options.Options,
		}, nil
	}

	return &db.Postgres{}, db.TypeError{Type: options.Type}
}

// Export exports given tables from provider.
func Export() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export provider table to local",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbType, _ := cmd.Flags().GetString("type")
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
				}

				b := backup.New(backupOptions)

				err = database.Export(&b)
				if err != nil {
					return err
				}

				fmt.Printf("Database %s exported successfully.\n", databaseName)
			}

			return nil
		},
	}

	SetupExportCommand(cmd)

	return cmd
}
