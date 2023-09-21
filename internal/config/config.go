package config

import (
	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
)

// Config represents struct for an app configuration.
type Config struct {
	Frontend *Frontend `koanf:"frontend"`
}

// Frontend represents struct for a frontend part of the config file.
type Frontend struct {
	HTMX        string `koanf:"htmx"`
	Hyperscript string `koanf:"hyperscript"`
	CSS         *CSS   `koanf:"css"`
}

// CSS represents struct for a css frontend part of the config file.
type CSS struct {
	Framework string `koanf:"framework"`
	Version   string `koanf:"version"`
}

// New creates a new config.
func New() (*Config, error) {
	// Parse a default YAML config file from the given path to a struct.
	return helpers.ParseYAMLToStruct(constants.YAMLConfigFileName, &Config{})
}
