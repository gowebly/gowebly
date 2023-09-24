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
	Port    int      `koanf:"port"`
	Name    string   `koanf:"name"`
	Timeout *Timeout `koanf:"timeout"`
}

// Timeout represents struct for a 'timeout' block in the backend part of the config file.
type Timeout struct {
	Read  int `koanf:"read"`
	Write int `koanf:"write"`
}

// Frontend represents struct for a frontend part of the config file.
type Frontend struct {
	HTMX         string `koanf:"htmx"`
	Hyperscript  string `koanf:"hyperscript"`
	CSSFramework string `koanf:"css_framework"`
}

// New creates a new config.
func New() (*Config, error) {
	// Parse a default YAML config file from the given path to a struct.
	return helpers.ParseYAMLToStruct(constants.YAMLConfigFileName, &Config{})
}
