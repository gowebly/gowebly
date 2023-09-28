package validators

import (
	"errors"
	"slices"

	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/constants"
)

// Validate validates the given config.
func Validate(cfg *config.Config) error {
	// Check, if 'backend' block is set.
	if cfg.Backend == nil {
		return errors.New(constants.ErrorValidateConfigBackendBlockNotFound)
	}

	// Check, if 'backend' block is set.
	if cfg.Backend != nil {
		// Check, if 'module_name' option in the 'backend' block is not set.
		if cfg.Backend.ModuleName == "" {
			return errors.New(constants.ErrorValidateConfigBackendModuleNameNotFound)
		}

		// Check, if 'go_framework' option in the 'backend' block is not set.
		if cfg.Backend.GoFramework == "" {
			return errors.New(constants.ErrorValidateConfigBackendGoFrameworkNotFound)
		}

		// Check, if 'go_framework' option in the 'backend' block has a known value.
		if !slices.Contains([]string{"default", "fiber", "echo", "chi"}, cfg.Backend.GoFramework) {
			return errors.New(constants.ErrorValidateConfigBackendGoFrameworkUnknown)
		}

		// Check, if 'port' option in the 'backend' block is not set.
		if cfg.Backend.Port == 0 {
			return errors.New(constants.ErrorValidateConfigBackendPortNotFound)
		}

		// Check, if 'timeout' option in the 'backend' block is not set.
		if cfg.Backend.Timeout == nil {
			return errors.New(constants.ErrorValidateConfigBackendTimeoutNotFound)
		}
	}

	// Check, if the 'frontend' block is not set.
	if cfg.Frontend == nil {
		return errors.New(constants.ErrorValidateConfigFrontendBlockNotFound)
	}

	// Check, if the 'frontend' block is set.
	if cfg.Frontend != nil {
		// Check, if 'package_name' option in the 'frontend' block is not set.
		if cfg.Frontend.PackageName == "" {
			return errors.New(constants.ErrorValidateConfigFrontendPackageNameNotFound)
		}

		// Check, if 'css_framework' option in the 'frontend' block is not set.
		if cfg.Frontend.CSSFramework == "" {
			return errors.New(constants.ErrorValidateConfigFrontendCSSFrameworkNotFound)
		}

		// Check, if 'css_framework' option in the 'frontend' block has a known value.
		if !slices.Contains([]string{"default", "tailwindcss", "unocss"}, cfg.Frontend.CSSFramework) {
			return errors.New(constants.ErrorValidateConfigFrontendCSSFrameworkUnknown)
		}

		// Check, if 'runtime_environment' option in the 'frontend' block is not set.
		if cfg.Frontend.RuntimeEnvironment == "" {
			return errors.New(constants.ErrorValidateConfigFrontendRuntimeEnvironmentNotFound)
		}

		// Check, if 'runtime_environment' option in the 'frontend' block has a known value.
		if !slices.Contains([]string{"default", "bun"}, cfg.Frontend.RuntimeEnvironment) {
			return errors.New(constants.ErrorValidateConfigFrontendRuntimeEnvironmentUnknown)
		}

		// Check, if 'htmx' option in the 'frontend' block is not set.
		if cfg.Frontend.HTMX == "" {
			return errors.New(constants.ErrorValidateConfigFrontendHTMXNotFound)
		}

		// Check, if 'hyperscript' option in the 'frontend' block is not set.
		if cfg.Frontend.Hyperscript == "" {
			return errors.New(constants.ErrorValidateConfigFrontendHyperscriptNotFound)
		}
	}

	return nil
}
