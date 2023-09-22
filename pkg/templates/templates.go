package templates

import (
	"html/template"

	"github.com/gowebly/gowebly/internal/constants"
)

// AddTemplates adds the given templates to the HTTP handler.
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
//		tmpl := gowebly.AddTemplates(
//			"templates/pages/index.html",
//			"templates/components/login-form.html",
//		)
//
//		if err := tmpl.Execute(w, nil); err != nil {
//			slog.Error("msg", err.Error())
//			return
//		}
//	}
func AddTemplates(names ...string) *template.Template {
	// Set global templates.
	global := []string{
		constants.TemplateHTMLMainFileName,
		constants.TemplateHTMLStylesFileName,
		constants.TemplateHTMLScriptsFileName,
	}

	// Add all user templates after global.
	global = append(global, names...)

	return template.Must(template.ParseFiles(global...))
}
