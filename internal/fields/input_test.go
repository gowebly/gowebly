package fields

import (
	"testing"

	"github.com/gowebly/gowebly/v3/internal/config"
	"github.com/gowebly/gowebly/v3/internal/injectors"
	"github.com/stretchr/testify/require"
)

func TestGoModuleNameInput(t *testing.T) {
	di := &injectors.Injector{
		Config: &config.Config{
			Backend: &config.Backend{
				ModuleName: "github.com/user/project",
			},
		},
	}

	require.NoError(t, GoModuleNameInput(di).Error())
}

func TestPortInput(t *testing.T) {
	di := &injectors.Injector{
		Config: &config.Config{
			Backend: &config.Backend{
				Port: "7000",
			},
		},
	}

	require.NoError(t, PortInput(di).Error())
}

func TestPackageNameInput(t *testing.T) {
	di := &injectors.Injector{
		Config: &config.Config{
			Frontend: &config.Frontend{
				PackageName: "project",
			},
		},
	}

	require.NoError(t, PackageNameInput(di).Error())
}
