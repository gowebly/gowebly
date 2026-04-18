package config

import "runtime"

// Config holds the complete project configuration.
type Config struct {
	Backend    *Backend
	Frontend   *Frontend
	Tools      *Tools
	SystemInfo *SystemInfo
}

// Backend contains Go backend configuration options.
type Backend struct {
	ModuleName, GoFramework, Port string
}

// Frontend contains frontend stack configuration options.
type Frontend struct {
	PackageName, ReactivityLibrary, CSSFramework string
}

// Tools contains optional development tool selections.
type Tools struct {
	IsUseAir, IsUseBun, IsUseTempl, IsUseGolangCILint bool
}

// SystemInfo contains detected operating system information.
type SystemInfo struct {
	IsArch64, IsWindows, IsDarwin, IsLinux bool
}

// New returns a Config with sensible default values.
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
