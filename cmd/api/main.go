package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	var appConfig Config

	err := appConfig.LoadDotEnvConfig()

	if err := appConfig.LoadDotEnvConfig(); err != nil {
		slog.Error("Unable to load configuration from .env file. Verify file exists and is correct.")
		os.Exit(1)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &Application{
		config: appConfig,
	}

	server := &http.Server{
		Addr:     fmt.Sprintf("%s:%d", app.config.Host, app.config.Port),
		Handler:  app.Routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Musicpedia API is up and running.",
		"Host", app.config.Host,
		"Port", app.config.Port,
	)

	err = server.ListenAndServe()

	logger.Error(err.Error())

	os.Exit(-1)
}
