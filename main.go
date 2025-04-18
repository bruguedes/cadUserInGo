package main

import (
	"cadUser/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := structLogger()

	if err := run(); err != nil {
		logger.Error("Error occurred", "error", err)
		return

	}

	logger.Info("System online")

}

func run() error {

	handler := api.NewHandler()

	server := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := server.ListenAndServe(); err != nil {
		return err

	}

	return nil

}

func structLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
