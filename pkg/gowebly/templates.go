package gowebly

import (
	"html/template"

	"github.com/gowebly/gowebly/internal/constants"
)

// AddHTMLTemplates adds the given HTML templates to a handler.
//
// Example:
//
//	tmpl := gowebly.AddHTMLTemplates(
//		"templates/pages/index.html",
//		"templates/components/login-form.html",
//	)
//
//	if err := tmpl.Execute(w, nil); err != nil {
//		slog.Error("msg", err.Error())
//		return
//	}
func AddHTMLTemplates(names ...string) *template.Template {
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
