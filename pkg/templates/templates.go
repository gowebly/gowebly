package templates

import (
	"fmt"
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
		fmt.Sprintf("templates/%s", constants.TemplateHTMLMainFileName),
		fmt.Sprintf("templates/components/gowebly/%s", constants.TemplateHTMLStylesFileName),
		fmt.Sprintf("templates/components/gowebly/%s", constants.TemplateHTMLScriptsFileName),
	}

	// Add all user templates after global.
	global = append(global, names...)

	return template.Must(template.ParseFiles(global...))
}
