package config

import (
	"github.com/gowebly/gowebly/internal/helpers"
)

// Config represents struct for an app configuration.
type Config struct {
	Backend  *Backend  `koanf:"backend"`
	Frontend *Frontend `koanf:"frontend"`
}

// Backend represents struct for a backend part of the config file.
type Backend struct {
	ModuleName  string   `koanf:"module_name"`
	GoFramework string   `koanf:"go_framework"`
	Port        int      `koanf:"port"`
	Timeout     *Timeout `koanf:"timeout"`
}

// Timeout represents struct for a 'timeout' block in the backend part of the config file.
type Timeout struct {
	Read  int `koanf:"read"`
	Write int `koanf:"write"`
}

// Frontend represents struct for a frontend part of the config file.
type Frontend struct {
	PackageName        string    `koanf:"package_name"`
	CSSFramework       string    `koanf:"css_framework"`
	RuntimeEnvironment string    `koanf:"runtime_environment"`
	HTMX               string    `koanf:"htmx"`
	Hyperscript        string    `koanf:"hyperscript"`
	Manifest           *Manifest `koanf:"manifest"`
}

// Manifest represents struct for a 'manifest' block in the frontend part of the config file.
type Manifest struct {
	Name            string  `koanf:"name"`
	ShortName       string  `koanf:"short_name"`
	Description     string  `koanf:"description"`
	BackgroundColor string  `koanf:"background_color"`
	ThemeColor      string  `koanf:"theme_color"`
	Display         string  `koanf:"display"`
	Orientation     string  `koanf:"orientation"`
	StartURL        string  `koanf:"start_url"`
	Icons           []*Icon `koanf:"icons"`
}

// Icon represents struct for a 'icons' block in the frontend part of the config file.
type Icon struct {
	Src   string `koanf:"src"`
	Type  string `koanf:"type"`
	Sizes string `koanf:"sizes"`
}

// New creates a new config.
func New() (*Config, error) {
	// Parse a default YAML config file to a struct.
	return helpers.ParseYAMLToStruct(&Config{})
}
