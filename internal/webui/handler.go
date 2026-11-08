// Package webui serves the application UI embedded in the Go binary.
package webui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static/*
var assets embed.FS

// Handler serves only known UI files, leaving API and Swagger routing independent.
func Handler() http.Handler {
	files, err := fs.Sub(assets, "static")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(files))
}
