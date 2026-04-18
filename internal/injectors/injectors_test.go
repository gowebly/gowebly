package injectors

import (
	"testing"

	"github.com/gowebly/gowebly/v3/internal/attachments"
	"github.com/gowebly/gowebly/v3/internal/config"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	injector := New(&config.Config{Backend: &config.Backend{}, Frontend: &config.Frontend{}}, &attachments.Attachments{})
	require.NotNil(t, injector)
	require.NotNil(t, injector.Config)
	require.NotNil(t, injector.Attachments)
}
