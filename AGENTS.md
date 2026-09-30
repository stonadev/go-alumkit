# AGENTS.md

Go module: `github.com/stonadev/alumkit`. Library consumed by Go web apps.

## Structure

```
alumkit.go          # Public API: New(), App, Config, LoadConfig()
internal/
  config/           # Config struct + Load() from env/dotenv
  auth/             # Session store, RequireAuth middleware
  rbac/             # Casbin enforcer, RequirePermission middleware
  handler/          # Dashboard route handlers (/dashboard/*)
  middleware/       # CORS, CSP, security headers
  mail/             # SMTP sender
  db/               # Connection wrapper
  activitylog/      # Audit logging
example/            # Demo app (separate go.mod, uses replace directive)
resources/          # Source CSS (Tailwind v4), JS (Editor.js, Cropper.js)
static/             # Build output - embedded in binary via go:embed
```

## Key patterns

- `New(cfg *Config) *App` - DB connection mandatory, fatal on failure
- Admin routes auto-mounted at `/dashboard` via chi `Mount()`
- Config: env vars only, no defaults for required fields (DB_*, SESSION_KEY)
- Types re-exported: `alumkit.Config = config.Config`
- Static files embedded via `//go:embed static/*` (CSS, JS, sortable)
- Consuming apps mount their own assets: `app.MountStatic("/assets", fs)`

## Conventions

- Internal packages stay internal
- Handlers return `http.HandlerFunc`
- Tests: table-driven, `t.Parallel()` where safe, black-box (`package foo_test`)

## Gotchas

- `New()` calls `log.Fatal` on missing config - test via subprocess if needed
- `example/go.mod` uses `replace` directive for local dev
- `t.Setenv` cannot use `t.Parallel` (Go restriction)
