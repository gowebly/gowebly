package actions

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
)

// CreateBackendFiles creates backend files.
func CreateBackendFiles(di *injectors.Injector) error {
	// Define embed templates.
	templates := []helpers.EmbedTemplate{
		// Backend files.
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/go.mod.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "go.mod",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/server.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "server.go",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/handlers.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "handlers.go",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/main.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "main.go",
			Data:       nil,
		},
		// Deploy files.
		{
			EmbedFile:  "templates/deploy/dockerignore.gotmpl",
			OutputFile: ".dockerignore",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  "templates/deploy/Dockerfile.gotmpl",
			OutputFile: "Dockerfile",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  "templates/deploy/docker-compose.yml.gotmpl",
			OutputFile: "docker-compose.yml",
			Data:       di.Config.Backend,
		},
		// Misc files.
		{
			EmbedFile:  "templates/misc/gitignore.gotmpl",
			OutputFile: ".gitignore",
			Data:       nil,
		},
		{
			EmbedFile:  "templates/misc/air.toml.gotmpl",
			OutputFile: "air.toml",
			Data:       nil,
		},
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, templates)
}
