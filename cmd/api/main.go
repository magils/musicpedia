package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	var appConfig Config

	flag.StringVar(&appConfig.Host, "host", "", "API Host Address")
	flag.IntVar(&appConfig.Port, "port", 4001, "API Server Port")
	flag.StringVar(&appConfig.Env, "env", "dev", "App deployment environment")

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

	err := server.ListenAndServe()

	logger.Error(err.Error())

	os.Exit(-1)
}
