package main

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/neatflowcv/loa-key-pool/cmd/loa-key-pool/api"
)

//go:generate go tool oapi-codegen -config ../../docs/cfg.yaml ../../docs/openapi.json

func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	return info.Main.Version
}

func main() {
	log.Println("version", version())

	service := NewService()
	router := chi.NewMux()

	handler := api.HandlerFromMux(api.NewStrictHandler(service, nil), router)

	const timeout = 10 * time.Second

	server := &http.Server{ //nolint:exhaustruct
		ReadHeaderTimeout: timeout,
		Addr:              ":8080",
		Handler:           handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
