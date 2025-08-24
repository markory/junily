package assets

import (
	"embed"
)

//go:embed dashboard/*
var StaticFiles embed.FS
