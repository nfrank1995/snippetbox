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

    logLevelMap := make(map[string]slog.Level)
    logLevelMap["Debug"] = slog.LevelDebug
    logLevelMap["Info"] = slog.LevelInfo
    logLevelMap["Warning"] = slog.LevelWarn
    logLevelMap["Error"] = slog.LevelError

    slogLevel, logFlagOk := logLevelMap[*logLevel]

    logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
      Level: slogLevel,
    }))

    if !logFlagOk {
      logger.Warn("could not set log level, fall back to log level Info", slog.String("log-level", *logLevel))
    }

    
    app := &application{
        logger: logger,
    }

    mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))

    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet/view", app.snippetView)
    mux.HandleFunc("/snippet/create", app.snippetCreate)

    logger.Info("starting server", "addr", *addr)

    err := http.ListenAndServe(*addr, mux)
    logger.Error(err.Error())
    os.Exit(1)
}
