package fields

import (
	"testing"

	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/injectors"
	"github.com/gowebly/gowebly/internal/messages"
	"github.com/stretchr/testify/require"
)

func TestGoFrameworkSelect(t *testing.T) {
	di := &injectors.Injector{
		Config: &config.Config{
			Backend: &config.Backend{
				GoFramework: "chi",
			},
		},
	}

	require.NoError(t, GoFrameworkSelect(di).Error())
	require.EqualValues(t, GoFrameworkSelect(di).GetValue(), messages.FormGoFrameworkChiValue)
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
	require.EqualValues(t, CSSFrameworkSelect(di).GetValue(), messages.FormCSSFrameworkDefaultValue)
}
