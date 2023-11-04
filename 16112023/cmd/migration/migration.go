package migration

import (
	"context"
	"github.com/kitabisa/kibitalk/config"
	"github.com/kitabisa/kibitalk/config/database"
	plog "github.com/kitabisa/perkakas/log"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var MigrateUpCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Up",
	Run: func(cmd *cobra.Command, args []string) {
		config.NewAppConfig()
		database.InitMySQL()
		mSource := getMigrateSource()
		total, err := migrate.Exec(database.MySqlDB.GetDB(), "mysql", mSource, migrate.Up)
		if err != nil {
			log.Printf("Fail migration | %v", err)
			os.Exit(1)
		}

		plog.Zlogger(context.Background()).Info().Msgf("Migrate Success, total migrated: %d", total)
	},
}

func getMigrateSource() migrate.FileMigrationSource {
	source := migrate.FileMigrationSource{
		Dir: "migrations/sql",
	}

	return source
}
