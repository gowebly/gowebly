package validators

import (
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
	require.EqualError(t, err, constants.ErrorValidateConfigBackendBlockNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: nil,
	})
	require.EqualError(t, err, constants.ErrorValidateConfigBackendModuleNameNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: nil,
	})
	require.EqualError(t, err, constants.ErrorValidateConfigBackendGoFrameworkNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "unknown", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: nil,
	})
	require.EqualError(t, err, constants.ErrorValidateConfigBackendGoFrameworkUnknown)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 0, Timeout: &config.Timeout{}},
		Frontend: nil,
	})
	require.EqualError(t, err, constants.ErrorValidateConfigBackendPortNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: nil},
		Frontend: nil,
	})
	require.EqualError(t, err, constants.ErrorValidateConfigBackendTimeoutNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: nil,
	})
	require.EqualError(t, err, constants.ErrorValidateConfigFrontendBlockNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: &config.Frontend{PackageName: "", CSSFramework: "", HTMX: "", Hyperscript: ""},
	})
	require.EqualError(t, err, constants.ErrorValidateConfigFrontendPackageNameNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: &config.Frontend{PackageName: "project", CSSFramework: "", HTMX: "", Hyperscript: ""},
	})
	require.EqualError(t, err, constants.ErrorValidateConfigFrontendCSSFrameworkNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: &config.Frontend{PackageName: "project", CSSFramework: "", HTMX: "", Hyperscript: ""},
	})
	require.EqualError(t, err, constants.ErrorValidateConfigFrontendCSSFrameworkNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: &config.Frontend{PackageName: "project", CSSFramework: "unknown", HTMX: "", Hyperscript: ""},
	})
	require.EqualError(t, err, constants.ErrorValidateConfigFrontendCSSFrameworkUnknown)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: &config.Frontend{PackageName: "project", CSSFramework: "default", HTMX: "", Hyperscript: ""},
	})
	require.EqualError(t, err, constants.ErrorValidateConfigFrontendHTMXNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: &config.Frontend{PackageName: "project", CSSFramework: "default", HTMX: "latest", Hyperscript: ""},
	})
	require.EqualError(t, err, constants.ErrorValidateConfigFrontendHyperscriptNotFound)

	err = Validate(&config.Config{
		Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{}},
		Frontend: &config.Frontend{PackageName: "project", CSSFramework: "default", HTMX: "latest", Hyperscript: "latest"},
	})
	require.NoError(t, err)
}
