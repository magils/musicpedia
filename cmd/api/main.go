package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"com.mgil.musicpedia/internal/db"
	"com.mgil.musicpedia/internal/models"
)

func main() {
	var appConfig Config

	err := appConfig.LoadDotEnvConfig()
	appConfig.ShowConfigurationValues()

	if err := appConfig.LoadDotEnvConfig(); err != nil {
		slog.Error("Unable to load configuration from .env file. Verify file exists and is correct.")
		os.Exit(1)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	dsn := appConfig.GetDsn()
	db, err := db.CreateDBConnection(dsn, 5, 5, 10*time.Minute)

	if err != nil {
		fmt.Println("DB connection error") // TODO: Improve this error handling and logging
		fmt.Println(err)
		os.Exit(1)
	}

	app := &Application{
		config:       appConfig,
		repositories: models.NewRepositories(db),
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
