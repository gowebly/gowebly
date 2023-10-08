package attachments

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	require.EqualValues(
		t,
		New(),
		&Files{Configs: ConfigsFiles, Templates: TemplatesFiles, Static: StaticFiles},
	)
}
