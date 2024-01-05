package fields

import (
	"testing"

	"github.com/gowebly/gowebly/v2/internal/config"
	"github.com/gowebly/gowebly/v2/internal/injectors"
	"github.com/stretchr/testify/require"
)

func TestIsUseAirConfirm(t *testing.T) {
	di := &injectors.Injector{
		Config: &config.Config{
			Tools: &config.Tools{
				IsUseAir: true,
			},
		},
	}

	require.NoError(t, IsUseAirConfirm(di).Error())
}

func TestIsUseTempleConfirm(t *testing.T) {
	di := &injectors.Injector{
		Config: &config.Config{
			Tools: &config.Tools{
				IsUseTempl: true,
			},
		},
	}

	require.NoError(t, IsUseTempleConfirm(di).Error())
}

func TestIsUseBunConfirm(t *testing.T) {
	di := &injectors.Injector{
		Config: &config.Config{
			Tools: &config.Tools{
				IsUseBun: true,
			},
		},
	}

	require.NoError(t, IsUseBunConfirm(di).Error())
}
