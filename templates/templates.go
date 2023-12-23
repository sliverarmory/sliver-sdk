package templates

import (
	"embed"
)

//go:embed extensions/go/**
var GoTemplates embed.FS

//go:embed extensions/rust/**
var RustTemplates embed.FS

//go:embed traffic_encoders/rust/**
var RustTrafficEncoderTemplates embed.FS
