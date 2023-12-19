package templates

import (
	"embed"
)

//go:embed go/**
var GoTemplates embed.FS

//go:embed rust/**
var RustTemplates embed.FS
