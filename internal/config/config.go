package config

// Config represents struct for an app configuration.
type Config struct {
	Backend  *Backend
	Frontend *Frontend
	Tools    *Tools
}

// Backend represents struct for a backend part of the project config.
type Backend struct {
	ModuleName, GoFramework, Port string
}

// Frontend represents struct for a frontend part of the project config.
type Frontend struct {
	IsUseHTMX, IsUseHyperscript                bool
	PackageName, ReactiveLibrary, CSSFramework string
}

// Tools represents struct for a tools part of the project config.
type Tools struct {
	IsUseAir, IsUseBun, IsUseTempl bool
}

// New creates a new config.
func New() *Config {
	return &Config{
		Backend: &Backend{
			ModuleName: "github.com/user/project",
			Port:       "7000",
		},
		Frontend: &Frontend{
			IsUseHTMX:       true,
			ReactiveLibrary: "htmx",
			PackageName:     "project",
		},
		Tools: &Tools{
			IsUseAir: true,
			IsUseBun: true,
		},
	}
}
