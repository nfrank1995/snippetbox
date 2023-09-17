
package main

import (
  "log/slog"
  "os"
)

func createLogger(logLevel string) *slog.Logger {

    logLevelMap := make(map[string]slog.Level)
    logLevelMap["Debug"] = slog.LevelDebug
    logLevelMap["Info"] = slog.LevelInfo
    logLevelMap["Warning"] = slog.LevelWarn
    logLevelMap["Error"] = slog.LevelError

    slogLevel, logFlagOk := logLevelMap[logLevel]

    logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
      Level: slogLevel,
    }))

    if !logFlagOk {
      logger.Warn("could not set log level, fall back to log level Info", slog.String("log-level", logLevel))
    }

    return logger
}


    
