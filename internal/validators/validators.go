package validators

import (
	"errors"

	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/constants"
)

// Validate validates the given config.
func Validate(config *config.Config) error {
	// Check, if 'backend' block is set.
	if config.Backend == nil {
		return errors.New(constants.ErrorValidateConfigBackendBlockNotFound)
	}

	// Check, if 'backend' block is set.
	if config.Backend != nil {
		// Check, if 'name' option in the 'backend' block is not set.
		if config.Backend.Name == "" {
			return errors.New(constants.ErrorValidateConfigBackendNameNotFound)
		}

		// Check, if 'port' option in the 'backend' block is not set.
		if config.Backend.Port == 0 {
			return errors.New(constants.ErrorValidateConfigBackendPortNotFound)
		}
	}

	// Check, if the 'frontend' block is not set.
	if config.Frontend == nil {
		return errors.New(constants.ErrorValidateConfigFrontendBlockNotFound)
	}

	// Check, if the 'frontend' block is set.
	if config.Frontend != nil {
		// Check, if 'htmx' option in the 'frontend' block is not set.
		if config.Frontend.HTMX == "" {
			return errors.New(constants.ErrorValidateConfigFrontendHTMXNotFound)
		}

		// Check, if 'hyperscript' option in the 'frontend' block is not set.
		if config.Frontend.Hyperscript == "" {
			return errors.New(constants.ErrorValidateConfigFrontendHyperscriptNotFound)
		}
	}

	return nil
}
