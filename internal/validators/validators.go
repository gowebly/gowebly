package validators

import (
	"fmt"
	"slices"

	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/constants"
)

// Validate validates the given config.
func Validate(cfg *config.Config) error {
	// Check, if 'backend' block is set.
	if cfg.Backend == nil {
		return fmt.Errorf(constants.ErrorValidateConfigBlockNotFound, "backend")
	}

	// Check, if 'backend' block is set.
	if cfg.Backend != nil {
		// Check, if 'module_name' option in the 'backend' block is not set.
		if cfg.Backend.ModuleName == "" {
			return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "module_name", "backend")
		}

		// Check, if 'go_framework' option in the 'backend' block is not set.
		if cfg.Backend.GoFramework == "" {
			return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "go_framework", "backend")
		}

		// Check, if 'go_framework' option in the 'backend' block has a known value.
		if !slices.Contains([]string{"default", "fiber", "echo", "chi"}, cfg.Backend.GoFramework) {
			return fmt.Errorf(
				constants.ErrorValidateConfigValueInOptionUnknown,
				"go_framework", "backend", cfg.Backend.GoFramework,
			)
		}

		// Check, if 'port' option in the 'backend' block is not set.
		if cfg.Backend.Port == 0 {
			return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "port", "backend")
		}

		// Check, if 'timeout' option in the 'backend' block is not set.
		if cfg.Backend.Timeout == nil {
			return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "timeout", "backend")
		}
	}

	// Check, if the 'frontend' block is not set.
	if cfg.Frontend == nil {
		return fmt.Errorf(constants.ErrorValidateConfigBlockNotFound, "frontend")
	}

	// Check, if the 'frontend' block is set.
	if cfg.Frontend != nil {
		// Check, if 'package_name' option in the 'frontend' block is not set.
		if cfg.Frontend.PackageName == "" {
			return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "package_name", "frontend")
		}

		// Check, if 'css_framework' option in the 'frontend' block is not set.
		if cfg.Frontend.CSSFramework == "" {
			return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "css_framework", "frontend")
		}

		// Check, if 'css_framework' option in the 'frontend' block has a known value.
		if !slices.Contains([]string{"default", "tailwindcss", "unocss"}, cfg.Frontend.CSSFramework) {
			return fmt.Errorf(
				constants.ErrorValidateConfigValueInOptionUnknown,
				"css_framework", "frontend", cfg.Frontend.CSSFramework,
			)
		}

		// Check, if 'runtime_environment' option in the 'frontend' block is not set.
		if cfg.Frontend.RuntimeEnvironment == "" {
			return fmt.Errorf(
				constants.ErrorValidateConfigOptionInBlockNotFound,
				"runtime_environment", "frontend",
			)
		}

		// Check, if 'runtime_environment' option in the 'frontend' block has a known value.
		if !slices.Contains([]string{"default", "bun"}, cfg.Frontend.RuntimeEnvironment) {
			return fmt.Errorf(
				constants.ErrorValidateConfigValueInOptionUnknown,
				"runtime_environment", "frontend", cfg.Frontend.RuntimeEnvironment,
			)
		}

		// Check, if 'htmx' option in the 'frontend' block is not set.
		if cfg.Frontend.HTMX == "" {
			return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "htmx", "frontend")
		}

		// Check, if 'hyperscript' option in the 'frontend' block is not set.
		if cfg.Frontend.Hyperscript == "" {
			return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "hyperscript", "frontend")
		}
	}

	return nil
}
