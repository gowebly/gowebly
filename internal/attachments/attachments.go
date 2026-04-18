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

// Attachments holds the embedded static files and templates.
type Attachments struct {
	Static, Templates embed.FS
}

// New returns Attachments with the embedded filesystems.
func New() *Attachments {
	return &Attachments{
		Static:    StaticFiles,
		Templates: TemplatesFiles,
	}
}
