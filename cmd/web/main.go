package main

import (
  "database/sql"
  "flag"
  "log/slog"
  "net/http"
  "os"

  _ "github.com/go-sql-driver/mysql"
)

type application struct {
  logger *slog.Logger
}

func main() {
    addr := flag.String("addr", ":4000", "HTTP network address")
    logLevel := flag.String("log-level", "Info", "The minimum log level to log. Supported severity levels are Debug, Info, Warning and Error in that order.")
    dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

    flag.Parse()

    logger := createLogger(*logLevel)

    db, err := openDB(*dsn)
    if err != nil {
      logger.Error(err.Error())
    }

    defer db.Close()

    app := &application{
        logger: logger,
    }

    logger.Info("starting server", "addr", *addr)

    err = http.ListenAndServe(*addr, app.routes())
    logger.Error(err.Error())
    os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }
  if err = db.Ping(); err != nil {
    return nil, err
  }
  return db, nil
}
