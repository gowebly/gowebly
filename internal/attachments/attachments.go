package attachments

import (
	"embed"
)

var (
	//go:embed configs/*
	ConfigsFiles embed.FS

	//go:embed templates/*
	TemplatesFiles embed.FS
)

// Files represents struct for embed files.
type Files struct {
	Configs, Templates embed.FS
}

// New creates a new collection with embed files by Files struct.
func New() *Files {
	return &Files{
		Configs:   ConfigsFiles,
		Templates: TemplatesFiles,
	}
}
