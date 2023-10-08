package injectors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gowebly/gowebly/internal/attachments"
	"github.com/gowebly/gowebly/internal/config"
)

func TestNew(t *testing.T) {
	_, err := New(
		&config.Config{Backend: nil, Frontend: nil},
		&attachments.Files{
			Configs:   attachments.ConfigsFiles,
			Templates: attachments.TemplatesFiles,
			Static:    attachments.StaticFiles,
		},
	)
	require.NoError(t, err)
}
