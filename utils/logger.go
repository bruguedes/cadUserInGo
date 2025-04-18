package utils

import (
	"log/slog"
	"os"
)

func StructLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
