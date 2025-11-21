package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
	}

	api := &application{
		config: cfg,
	}

	// handler_http := api.mount()
	// err:= api.run(handler_http)

	// if err!=nil{

	// }

	// strucutred logging. By using slog package
	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // we can also use json handler for slog.NewJsonHandler
	slog.SetDefault(logger)
	// passing this as a dependency. So all files make logs to this logger
	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
