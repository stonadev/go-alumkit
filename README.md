# AlumKit

A Go library for building alumni association websites with admin panel, content management, and static file handling.

## Installation

```bash
go get github.com/stonadev/alumkit
```

## Quick Start

```go
package main

import (
    "log"
    "net/http"
    "os"
    "github.com/stonadev/alumkit"
)

func main() {
    cfg, _ := alumkit.LoadConfig(".env")
    app := alumkit.New(cfg)
    
    // Mount your own static assets
    app.MountStatic("/assets", os.DirFS("public"))
    
    app.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, AlumKit!"))
    })
    
    log.Fatal(http.ListenAndServe(":8080", app))
}
```

## Configuration

Create a `.env` file:

```env
# Required
DB_HOST=localhost
DB_USER=alumkit
DB_PASSWORD=password
DB_NAME=alumkit
SESSION_KEY=your-secret-key-32-chars-min

# Optional
DB_PORT=5432
DB_SSLMODE=disable
APP_URL=http://localhost:8080
PORT=8080
```

## Static Assets

AlumKit embeds its admin panel assets automatically. For your own pages:

```go
// Serve your assets at /assets/*
app.MountStatic("/assets", os.DirFS("public"))
```

```html
<!-- In your templates -->
<link rel="stylesheet" href="/assets/app.css"/>
<script src="/assets/app.js"></script>
```

Build your assets:

```bash
# Copy package.json.example to your project
cp example/package.json.example package.json
npm install
npm run build
```

## Dashboard

Admin routes are auto-mounted at `/dashboard`. Login at `/dashboard/login`.

## License

MIT
