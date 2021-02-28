package dump

import (
	db "github.com/omegion/go-db-backup/pkg/database"
	"github.com/omegion/go-db-backup/pkg/store"
	"github.com/spf13/cobra"
	"log"
)

func setupDumpCommand(cmd *cobra.Command) {
	cmd.Flags().String("type", "postgres", "Database type")
	cmd.Flags().String("host", "", "Host")
	if err := cmd.MarkFlagRequired("host"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("port", "", "Port")
	if err := cmd.MarkFlagRequired("port"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("database", "", "Database name")
	if err := cmd.MarkFlagRequired("database"); err != nil {
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

func getDatabaseByType(options db.Options) (db.Database, error) {
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

func Local() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "local",
		Short: "Print the version/build number",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbType, _ := cmd.Flags().GetString("type")
			host, _ := cmd.Flags().GetString("host")
			port, _ := cmd.Flags().GetString("port")
			databaseName, _ := cmd.Flags().GetString("database")
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")

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

			export := store.Local{}

			err = result.To(&export)
			if err != nil {
				return err
			}

			return nil
		},
	}

	setupDumpCommand(cmd)

	return cmd
}
