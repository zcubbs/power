package docs

import "embed"

//go:embed swagger/*
var SwaggerDist embed.FS
