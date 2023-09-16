package main

import (
  "flag"
  "log/slog"
  "net/http"
  "os"
)

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

    mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))

    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

    logger.Info("starting server", "addr", *addr)

    err := http.ListenAndServe(*addr, mux)
    logger.Error(err.Error())
    os.Exit(1)
}
