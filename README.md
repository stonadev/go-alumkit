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
    "github.com/stonadev/alumkit"
)

func main() {
    cfg, _ := alumkit.LoadConfig(".env")
    app := alumkit.New(*cfg)
    
    app.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, AlumKit!"))
    })
    
    log.Fatal(http.ListenAndServe(":8080", app))
}
```

## Configuration

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=alumkit
DB_PASSWORD=password
DB_NAME=alumkit
SESSION_KEY=your-secret-key-32-chars-min
APP_URL=http://localhost:8080
```

## Building Assets

```bash
npm install
npm run build:css
npm run build:js
```

## License

MIT
