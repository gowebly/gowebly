package config

import "runtime"

// Config represents struct for an app configuration.
type Config struct {
	Backend    *Backend
	Frontend   *Frontend
	Tools      *Tools
	SystemInfo *SystemInfo
}

// Backend represents struct for a backend part of the project config.
type Backend struct {
	ModuleName, GoFramework, Port string
}

// Frontend represents struct for a frontend part of the project config.
type Frontend struct {
	PackageName, ReactivityLibrary, CSSFramework string
}

// Tools represents struct for a tools part of the project config.
type Tools struct {
	IsUseAir, IsUseBun, IsUseTempl, IsUseGolangCILint bool
}

// SystemInfo represents struct for a system info part of the project config.
type SystemInfo struct {
	IsArch64, IsWindows, IsDarwin, IsLinux bool
}

// New creates a new config.
func New() *Config {
	return &Config{
		Backend: &Backend{
			ModuleName:  "github.com/user/project",
			GoFramework: "default",
			Port:        "7000",
		},
		Frontend: &Frontend{
			PackageName:       "project",
			ReactivityLibrary: "htmx",
			CSSFramework:      "default",
		},
		Tools: &Tools{
			IsUseAir:          true,
			IsUseBun:          true,
			IsUseTempl:        true,
			IsUseGolangCILint: true,
		},
		SystemInfo: &SystemInfo{
			IsArch64:  runtime.GOARCH == "amd64",
			IsWindows: runtime.GOOS == "windows",
			IsDarwin:  runtime.GOOS == "darwin",
			IsLinux:   runtime.GOOS == "linux",
		},
	}
}
