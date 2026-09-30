package alumkit

import (
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// migrationsFS holds the SQL migration files shipped with the library.
//
//go:embed db/migrations/*.sql
var migrationsFS embed.FS

// migrationsDir is the path inside migrationsFS where migration files live.
const migrationsDir = "db/migrations"

// migrateCmd returns the cobra command that runs AlumKit database migrations.
// The dbURL must be a postgres:// connection string (see Config.DSN).
func migrateCmd(dbURL string) *cobra.Command {
	var fromVersion uint
	var toVersion uint

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Run AlumKit database migrations",
		Long: `Run AlumKit database migrations.

With no flags, applies all pending migrations (use on a fresh database).
Use --to 1 on an existing production database that already has the schema:
the idempotent baseline is a no-op and golang-migrate records it as applied.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			src, err := iofs.New(migrationsFS, migrationsDir)
			if err != nil {
				return fmt.Errorf("failed to load embedded migrations: %w", err)
			}

			m, err := migrate.NewWithSourceInstance("iofs", src, dbURL)
			if err != nil {
				return fmt.Errorf("failed to create migrator: %w", err)
			}
			defer m.Close()

			var migrateErr error
			switch {
			case toVersion > 0:
				migrateErr = m.Migrate(toVersion)
			case fromVersion > 0:
				migrateErr = m.Migrate(fromVersion)
			default:
				migrateErr = m.Up()
			}

			if errors.Is(migrateErr, migrate.ErrNoChange) {
				cmd.Println("No migration needed, database is up to date")
				return nil
			}
			if migrateErr != nil {
				return fmt.Errorf("migration failed: %w", migrateErr)
			}

			version, dirty, verr := m.Version()
			if verr != nil && !errors.Is(verr, migrate.ErrNilVersion) {
				return verr
			}
			cmd.Printf("Migrated to version %d\n", version)
			if dirty {
				return errors.New("database is in dirty state; fix manually or reset")
			}
			return nil
		},
	}

	cmd.Flags().UintVar(&fromVersion, "from", 0, "Start from version")
	cmd.Flags().UintVar(&toVersion, "to", 0, "Migrate to specific version")

	return cmd
}
