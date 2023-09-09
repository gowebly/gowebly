package config

import (
	"github.com/gowebly/gowebly/internal/helpers"
)

// Config represents struct for an app configuration.
type Config struct {
	Frontend *Frontend `koanf:"frontend"`
	Storage  *Storage  `koanf:"storage"`
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

// Storage represents struct for a storage part of the config file.
type Storage struct {
	Postgres *Postgres `koanf:"postgres"`
}

// Postgres represents struct for a PostgreSQL storage part of the config file.
type Postgres struct {
	SelfHosted *SelfHosted `koanf:"self_hosted"`
	Cloud      *Cloud      `koanf:"cloud"`
}

// SelfHosted represents struct for a self-hosted storage part of the config file.
type SelfHosted struct {
	Port          int    `koanf:"port"`
	ContainerName string `koanf:"container_name"`
	Version       string `koanf:"version"`
	User          string `koanf:"user"`
	Password      string `koanf:"password"`
	Database      string `koanf:"database"`
}

// Cloud represents struct for a cloud storage part of the config file.
type Cloud struct {
	Port     int    `koanf:"port"`
	Host     string `koanf:"host"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Database string `koanf:"database"`
}

// New creates a new config from the given path.
func New(path string) (*Config, error) {
	// Parse the given YAML config file from the given path to a struct.
	return helpers.ParseYAMLToStruct(path, &Config{})
}
