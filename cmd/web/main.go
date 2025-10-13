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

type config struct {
	addr      string
	staticDir string
}

var cfg config

func main() {
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	app := &application{
		logger: logger,
	}

	// Example where handlers are in seperate package and need dependency injection
	// app = exampleconfig.Application{
	// 	Logger: logger,
	// }

	// mux.HandleFunc("/", foo.ExampleHandler(app))

	logger.Info("Starting server", slog.String("addr", cfg.addr))

	if err := http.ListenAndServe(cfg.addr, app.routes()); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
