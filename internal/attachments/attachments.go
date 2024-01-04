package attachments

import (
	"embed"
)

var (
	//go:embed all:templates
	TemplatesFiles embed.FS

	//go:embed all:static
	StaticFiles embed.FS
)

// Attachments represents struct for embed files.
type Attachments struct {
	Templates, Static embed.FS
}

// New creates a new collection with embed files by Attachments struct.
func New() *Attachments {
	return &Attachments{
		Templates: TemplatesFiles,
		Static:    StaticFiles,
	}
}
