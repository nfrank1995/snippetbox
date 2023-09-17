package main

import (
  "flag"
  "log/slog"
  "net/http"
  "os"
)

type application struct {
  logger *slog.Logger
}

func main() {
    addr := flag.String("addr", ":4000", "HTTP network address")
    logLevel := flag.String("log-level", "Info", "The minimum log level to log. Supported severity levels are Debug, Info, Warning and Error in that order.")

    flag.Parse()


    app := &application{
        logger: createLogger(*logLevel),
    }

    app.logger.Info("starting server", "addr", *addr)

    err := http.ListenAndServe(*addr, app.routes())
    app.logger.Error(err.Error())
    os.Exit(1)
}
