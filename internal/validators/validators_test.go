package validators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/constants"
)

func TestValidate(t *testing.T) {
	err := Validate(&config.Config{
		Backend:  nil,
		Frontend: nil,
	})
	require.EqualError(t, err, fmt.Sprintf(constants.ErrorValidateConfigBlockNotFound, "backend"))

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: nil},
		Frontend: nil,
	})
	require.EqualError(t, err, fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "timeout", "backend"))

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 0, Write: 0}},
		Frontend: nil,
	})
	require.EqualError(t, err, fmt.Sprintf(constants.ErrorValidateConfigBlockNotFound, "frontend"))

	timeout := &config.Timeout{Read: 0, Write: 0}

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: timeout},
		Frontend: &config.Frontend{PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "default", HTMX: "latest", Hyperscript: "latest"},
	})
	require.NoError(t, err)
}
