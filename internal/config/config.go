package config

import (
	"github.com/gowebly/gowebly/internal/constants"
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
	PackageName        string `koanf:"package_name"`
	CSSFramework       string `koanf:"css_framework"`
	RuntimeEnvironment string `koanf:"runtime_environment"`
	HTMX               string `koanf:"htmx"`
	Hyperscript        string `koanf:"hyperscript"`
}

// New creates a new config.
func New() (*Config, error) {
	// Parse a default YAML config file from the given path to a struct.
	return helpers.ParseYAMLToStruct(constants.YAMLConfigFileName, &Config{})
}
