package attachments

import (
	"embed"
)

var (
	//go:embed all:static
	StaticFiles embed.FS

	//go:embed all:templates
	TemplatesFiles embed.FS
)

// Attachments represents struct for embed files.
type Attachments struct {
	Static, Templates embed.FS
}

// New creates a new collection with embed files by Attachments struct.
func New() *Attachments {
	return &Attachments{
		Static:    StaticFiles,
		Templates: TemplatesFiles,
	}
}
