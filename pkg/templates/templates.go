package templates

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
)

// ParseTemplates parses list of the given templates to the HTTP handler.
//
// Already included 'templates/main.html' layout template from your project
// path.
//
// Example:
//
//	import (
//		"log/slog"
//
//		gowebly "github.com/gowebly/gowebly/pkg/templates"
//	)
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//		// Define paths to the user templates.
//		indexPage := filepath.Join("templates", "pages", "index.html")
//		indexLoginForm := filepath.Join("templates", "components", "index-login-form.html")
//
//		// Parse user templates or return error.
//		tmpl, err := gowebly.ParseTemplates(indexPage, indexLoginForm)
//		if err != nil {
//			slog.Error("msg", err.Error())
//			return
//		}
//
//		// Execute (render) all templates or return error.
//		if err := tmpl.Execute(w, nil); err != nil {
//			slog.Error("msg", err.Error())
//			return
//		}
//	}
func ParseTemplates(names ...string) (*template.Template, error) {
	// Set global templates.
	global := []string{
		filepath.Join("templates", constants.TemplateHTMLMainFileName),
	}

	for _, n := range names {
		// Check, if the given template is existing.
		if helpers.IsExistInFolder(n, false) {
			return nil, fmt.Errorf("os: template '%s' is not found", n)
		}
	}

	// Add all user templates after global.
	global = append(global, names...)

	return template.Must(template.ParseFiles(global...)), nil
}
