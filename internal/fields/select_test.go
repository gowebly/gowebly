package fields

import (
	"testing"

	"github.com/gowebly/gowebly/v2/internal/config"
	"github.com/gowebly/gowebly/v2/internal/injectors"
	"github.com/gowebly/gowebly/v2/internal/variables"
	"github.com/stretchr/testify/require"
)

func TestGoFrameworkSelect(t *testing.T) {
	di := &injectors.Injector{
		Config: &config.Config{
			Backend: &config.Backend{
				GoFramework: "default",
			},
		},
	}

	require.NoError(t, GoFrameworkSelect(di).Error())
	require.EqualValues(t, GoFrameworkSelect(di).GetValue(), variables.ListGoFrameworks["default"][0])
}

func TestReactivityLibrarySelect(t *testing.T) {
	di := &injectors.Injector{
		Config: &config.Config{
			Frontend: &config.Frontend{
				ReactivityLibrary: "htmx",
			},
		},
	}

	require.NoError(t, ReactivityLibrarySelect(di).Error())
	require.EqualValues(t, ReactivityLibrarySelect(di).GetValue(), variables.ListReactivityLibraries["htmx"][0])
}

func TestCSSFrameworkSelect(t *testing.T) {
	di := &injectors.Injector{
		Config: &config.Config{
			Frontend: &config.Frontend{
				CSSFramework: "default",
			},
		},
	}

	require.NoError(t, CSSFrameworkSelect(di).Error())
	require.EqualValues(t, CSSFrameworkSelect(di).GetValue(), variables.ListCSSFrameworks["default"][0])
}
