package actions

import (
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
)

// CreateFrontendFiles creates frontend files.
//
// It takes an instance of the Injector struct as a parameter.
// It returns an error if there is any issue during file generation.
func CreateFrontendFiles(di *injectors.Injector) error {
	// Check if the frontend is configured to use HTMX.
	if di.Config.Frontend.IsUseHTMX {
		// Generate HTMX files.
		if err := generateHTMXFiles(di); err != nil {
			return err
		}
	}

	// Check which CSS framework is configured for the frontend.
	switch di.Config.Frontend.CSSFramework {
	case "daisyui", "flowbite", "tailwindcss":
		// Generate Tailwind CSS files for supported frameworks.
		if err := generateTailwindCSSFiles(di); err != nil {
			return err
		}
	case "unocss":
		// Generate Uno CSS files for the UnoCSS framework.
		if err := generateUnoCSSFiles(di); err != nil {
			return err
		}
	}

	return nil
}

func generateHTMXFiles(di *injectors.Injector) error {
	// Define the files to be generated.
	files := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/htmx/styles.scss.gotmpl",
			OutputFile: "assets/styles.scss",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/htmx/scripts.js.gotmpl",
			OutputFile: "assets/scripts.js",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/htmx/package.json.gotmpl",
			OutputFile: "package.json",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/htmx/index.html.gotmpl",
			OutputFile: "templates/index.html",
			Data:       di.Config.Frontend,
		},
	}

	// Check if Templ is used.
	if di.Config.Tools.IsUseTempl {
		files = append(files,
			helpers.EmbedTemplate{
				EmbedFile:  "templates/htmx/main.templ.gotmpl",
				OutputFile: "templates/main.templ",
				Data:       di.Config.Frontend,
			},
			helpers.EmbedTemplate{
				EmbedFile:  "templates/htmx/index.templ.gotmpl",
				OutputFile: "templates/pages/index.templ",
				Data:       di.Config.Frontend,
			},
		)
	} else {
		files = append(files,
			helpers.EmbedTemplate{
				EmbedFile:  "templates/htmx/main.html.gotmpl",
				OutputFile: "templates/main.html",
				Data:       di.Config.Frontend,
			},
			helpers.EmbedTemplate{
				EmbedFile:  "templates/htmx/index.html.gotmpl",
				OutputFile: "templates/pages/index.html",
				Data:       di.Config.Frontend,
			},
		)
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, files)
}

func generateTailwindCSSFiles(di *injectors.Injector) error {
	// Define the files to be generated.
	files := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/css/tailwind.config.js.gotmpl",
			OutputFile: "tailwind.config.js",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/css/postcssrc.gotmpl",
			OutputFile: ".postcssrc",
			Data:       di.Config.Frontend,
		},
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, files)
}

func generateUnoCSSFiles(di *injectors.Injector) error {
	// Define the files to be generated.
	files := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/css/uno.config.ts.gotmpl",
			OutputFile: "uno.config.ts",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/css/postcssrc.gotmpl",
			OutputFile: ".postcssrc",
			Data:       di.Config.Frontend,
		},
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, files)
}
