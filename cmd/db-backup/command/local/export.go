package local

import (
	"fmt"
	db "github.com/omegion/go-db-backup/pkg/database"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

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

func GetDatabaseByType(options db.Options) (db.Database, error) {
	if options.Type == "postgres" {
		return &db.Postgres{
			Host:     options.Host,
			Port:     options.Port,
			Name:     options.Name,
			Username: options.Username,
			Password: options.Password,
			Options:  options.Options,
		}, nil
	} else {
		return &db.Postgres{}, db.TypeError{Type: options.Type}
	}
}

func Export() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export database table to local",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbType, _ := cmd.Flags().GetString("type")
			host, _ := cmd.Flags().GetString("host")
			port, _ := cmd.Flags().GetString("port")
			databases, _ := cmd.Flags().GetString("databases")
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")

			for _, databaseName := range strings.Split(databases, ",") {
				options := db.Options{
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

				_, err = database.Export()
				if err != nil {
					return err
				}

				fmt.Print(fmt.Sprintf("Database %s exported successfully.\n", databaseName))
			}

			return nil
		},
	}

	SetupExportCommand(cmd)

	return cmd
}
