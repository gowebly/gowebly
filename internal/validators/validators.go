package validators

import (
	"fmt"
	"slices"
	"strings"

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

		// Check, if 'manifest' option in the 'frontend' block is set.
		if cfg.Frontend.Manifest != nil {
			// Check, if 'name' option in the 'manifest' block is not set.
			if cfg.Frontend.Manifest.Name == "" {
				return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "name", "manifest")
			}

			// Check, if 'short_name' option in the 'manifest' block is not set.
			if cfg.Frontend.Manifest.ShortName == "" {
				return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "short_name", "manifest")
			}

			// Check, if 'description' option in the 'manifest' block is not set.
			if cfg.Frontend.Manifest.Description == "" {
				return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "description", "manifest")
			}

			// Check, if 'background_color' option in the 'manifest' block is not set.
			if cfg.Frontend.Manifest.BackgroundColor == "" {
				return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "background_color", "manifest")
			}

			// Check, if 'background_color' option in the 'manifest' block has a wrong value.
			if !strings.HasPrefix(cfg.Frontend.Manifest.BackgroundColor, "#") {
				return fmt.Errorf(
					constants.ErrorValidateConfigValueInOptionInvalid,
					"background_color", "manifest", cfg.Frontend.Manifest.BackgroundColor,
				)
			}

			// Check, if 'theme_color' option in the 'manifest' block is not set.
			if cfg.Frontend.Manifest.ThemeColor == "" {
				return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "theme_color", "manifest")
			}

			// Check, if 'theme_color' option in the 'manifest' block has a wrong value.
			if !strings.HasPrefix(cfg.Frontend.Manifest.ThemeColor, "#") {
				return fmt.Errorf(
					constants.ErrorValidateConfigValueInOptionInvalid,
					"theme_color", "manifest", cfg.Frontend.Manifest.ThemeColor,
				)
			}

			// Check, if 'display' option in the 'manifest' block is not set.
			if cfg.Frontend.Manifest.Display == "" {
				return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "display", "manifest")
			}

			// Check, if 'display' option in the 'manifest' block has a wrong value.
			displayVariants := []string{"fullscreen", "standalone", "minimal-ui", "browser"}
			if !slices.Contains(displayVariants, cfg.Frontend.Manifest.Display) {
				return fmt.Errorf(
					constants.ErrorValidateConfigValueInOptionInvalid,
					"display", "manifest", cfg.Frontend.Manifest.Display,
				)
			}

			// Check, if 'orientation' option in the 'manifest' block is not set.
			if cfg.Frontend.Manifest.Orientation == "" {
				return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "orientation", "manifest")
			}

			// Check, if 'orientation' option in the 'manifest' block has a wrong value.
			orientationVariants := []string{
				"landscape", "landscape-primary", "landscape-secondary",
				"portrait", "portrait-primary", "portrait-secondary",
				"natural", "any",
			}
			if !slices.Contains(orientationVariants, cfg.Frontend.Manifest.Orientation) {
				return fmt.Errorf(
					constants.ErrorValidateConfigValueInOptionInvalid,
					"orientation", "manifest", cfg.Frontend.Manifest.Orientation,
				)
			}

			// Check, if 'start_url' option in the 'manifest' block is not set.
			if cfg.Frontend.Manifest.StartURL == "" {
				return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "start_url", "manifest")
			}

			// Check, if 'icons' option in the 'manifest' block is not set.
			if cfg.Frontend.Manifest.Icons == nil {
				return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "icons", "manifest")
			}

			// Check, if 'icons' option in the 'manifest' block is not empty.
			if len(cfg.Frontend.Manifest.Icons) == 0 {
				return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockEmpty, "icons", "manifest")
			}

			// Check, if 'icons' option in the 'manifest' block has values.
			if len(cfg.Frontend.Manifest.Icons) > 0 {
				// Check, if 'src' option in the 'icons' block is set.
				for _, icon := range cfg.Frontend.Manifest.Icons {
					// Check, if 'src' option in the 'icons' block is not set.
					if icon.Src == "" {
						return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "src", "icons")
					}

					// Check, if 'type' option in the 'icons' block is not set.
					if icon.Type == "" {
						return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "type", "icons")
					}

					// Check, if 'type' option in the 'icons' block has a wrong value.
					typeVariants := []string{"image/jpeg", "image/png", "image/svg+xml", "image/webp"}
					if !slices.Contains(typeVariants, icon.Type) {
						return fmt.Errorf(
							constants.ErrorValidateConfigValueInOptionInvalid,
							"type", "icons", icon.Type,
						)
					}

					// Check, if 'sizes' option in the 'icons' block is not set.
					if icon.Sizes == "" {
						return fmt.Errorf(constants.ErrorValidateConfigOptionInBlockNotFound, "sizes", "icons")
					}
				}
			}
		}
	}

	return nil
}
