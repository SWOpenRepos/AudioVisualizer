package server

import (
	"embed"
	"io/fs"
	"os"
)

const embedPath = "embed/public"

//go:embed embed/public/*
var embedFS embed.FS

var (
	publicFS fs.FS
	assetsFS fs.FS
)

func init() {
	var err error
	publicFS, err = fs.Sub(embedFS, embedPath)
	if err != nil {
		os.Exit(1)
	}
	assetsFS, err = fs.Sub(publicFS, "assets")
	if err != nil {
		os.Exit(1)
	}
}
