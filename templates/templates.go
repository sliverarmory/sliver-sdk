package templates

import (
	"embed"
)

//go:embed extensions/go/**
var GoExtensionTemplates embed.FS

//go:embed extensions/rust/**
var RustExtensionTemplates embed.FS

//go:embed encoders/rust/**
var RustTrafficEncoderTemplates embed.FS
