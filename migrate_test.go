package alumkit

import (
	"io/fs"
	"regexp"
	"strings"
	"testing"
)

// baselineTables lists the 23 tables exported from the PHP production
// database (alumkit-go) that the baseline migration must create.
var baselineTables = []string{
	"sessions",
	"permissions",
	"roles",
	"model_has_permissions",
	"model_has_roles",
	"role_has_permissions",
	"users",
	"password_reset_tokens",
	"profiles",
	"educations",
	"careers",
	"posts",
	"positions",
	"committee_members",
	"pages",
	"contents",
	"activity_log",
	"cache",
	"cache_locks",
	"failed_jobs",
	"job_batches",
	"jobs",
	"migrations",
}

func readMigration(t *testing.T, name string) string {
	t.Helper()
	data, err := migrationsFS.ReadFile(migrationsDir + "/" + name)
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return string(data)
}

func TestBaselineUpContainsAllTables(t *testing.T) {
	t.Parallel()
	sql := readMigration(t, "000001_baseline.up.sql")

	for _, table := range baselineTables {
		t.Run(table, func(t *testing.T) {
			t.Parallel()
			want := "CREATE TABLE IF NOT EXISTS public." + table + " ("
			if !strings.Contains(sql, want) {
				t.Errorf("baseline up migration missing %q", want)
			}
		})
	}
}

func TestBaselineDownDropsAllTables(t *testing.T) {
	t.Parallel()
	sql := readMigration(t, "000001_baseline.down.sql")

	for _, table := range baselineTables {
		t.Run(table, func(t *testing.T) {
			t.Parallel()
			want := "DROP TABLE IF EXISTS " + table + ";"
			if !strings.Contains(sql, want) {
				t.Errorf("baseline down migration missing %q", want)
			}
		})
	}
}

func TestBaselineUpIsIdempotent(t *testing.T) {
	t.Parallel()
	sql := readMigration(t, "000001_baseline.up.sql")

	// The baseline runs against the existing production database via
	// `migrate --to 1`, so it must be re-runnable: no role-specific OWNER
	// statements, and constraints guarded by IF NOT EXISTS checks.
	for _, banned := range []string{"OWNER TO", "GRANT ", "set_config"} {
		if strings.Contains(sql, banned) {
			t.Errorf("baseline up migration contains %q, which breaks portability/idempotency", banned)
		}
	}
	if got := strings.Count(sql, "IF NOT EXISTS (SELECT 1 FROM pg_constraint"); got == 0 {
		t.Error("baseline up migration has no guarded constraints")
	}
}

func TestSessionsSCSMigration(t *testing.T) {
	t.Parallel()

	// 000002 replaces the Laravel sessions table with the schema
	// github.com/alexedwards/scs/v2 postgresstore queries against.
	tests := []struct {
		file    string
		pattern string
	}{
		{"000002_sessions_scs.up.sql", `DROP TABLE IF EXISTS sessions`},
		{"000002_sessions_scs.up.sql", `token\s+text PRIMARY KEY`},
		{"000002_sessions_scs.up.sql", `data\s+bytea`},
		{"000002_sessions_scs.up.sql", `expiry\s+timestamptz NOT NULL`},
		{"000002_sessions_scs.up.sql", `sessions_expiry_idx`},
		{"000002_sessions_scs.down.sql", `payload\s+text NOT NULL`},
		{"000002_sessions_scs.down.sql", `last_activity\s+integer NOT NULL`},
		{"000002_sessions_scs.down.sql", `sessions_pkey`},
		{"000002_sessions_scs.down.sql", `sessions_last_activity_index`},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(tt.file+"/"+regexp.MustCompile(tt.pattern).String(), func(t *testing.T) {
			t.Parallel()
			sql := readMigration(t, tt.file)
			if !regexp.MustCompile(tt.pattern).MatchString(sql) {
				t.Errorf("case %d: %s missing /%s/", i, tt.file, tt.pattern)
			}
		})
	}
}

func TestMigrateCmd(t *testing.T) {
	t.Parallel()
	cmd := migrateCmd("postgres://user:pass@localhost:5432/alumkit?sslmode=disable")

	if cmd.Use != "migrate" {
		t.Errorf("Use = %q, want %q", cmd.Use, "migrate")
	}
	if cmd.Short == "" {
		t.Error("Short description is empty")
	}
	for _, flag := range []string{"from", "to"} {
		if cmd.Flags().Lookup(flag) == nil {
			t.Errorf("missing --%s flag", flag)
		}
	}
}

func TestEmbeddedMigrationsFS(t *testing.T) {
	t.Parallel()
	entries, err := fs.ReadDir(migrationsFS, migrationsDir)
	if err != nil {
		t.Fatalf("ReadDir(%s): %v", migrationsDir, err)
	}
	if len(entries) == 0 {
		t.Fatal("no embedded migration files")
	}
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".up.sql") && !strings.HasSuffix(e.Name(), ".down.sql") {
			t.Errorf("unexpected file %q (golang-migrate requires *.up.sql / *.down.sql)", e.Name())
		}
	}
}
