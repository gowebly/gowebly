package attachments

import (
	"embed"
)

var (
	//go:embed configs/*
	ConfigsFiles embed.FS

	//go:embed templates/*
	TemplatesFiles embed.FS

	//go:embed static/*
	StaticFiles embed.FS
)

// Files represents struct for embed files.
type Files struct {
	Configs, Templates, Static embed.FS
}

// New creates a new collection with embed files by Files struct.
func New() *Files {
	return &Files{
		Configs:   ConfigsFiles,
		Templates: TemplatesFiles,
		Static:    StaticFiles,
	}
}
