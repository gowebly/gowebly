package validators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/constants"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name    string
		config  *config.Config
		wantErr string
	}{
		{
			name: "Backend block not found",
			config: &config.Config{
				Backend:  nil,
				Frontend: nil,
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigBlockNotFound, "backend"),
		},
		{
			name: "Timeout option in Backend block not found",
			config: &config.Config{
				Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: nil},
				Frontend: nil,
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "timeout", "backend"),
		},
		{
			name: "ModuleName option in Backend block not found",
			config: &config.Config{
				Backend:  &config.Backend{ModuleName: "", GoFramework: "default", Port: 5000, Timeout: nil},
				Frontend: nil,
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "module_name", "backend"),
		},
		{
			name: "GoFramework option in Backend block not found",
			config: &config.Config{
				Backend:  &config.Backend{ModuleName: "project", GoFramework: "", Port: 5000, Timeout: nil},
				Frontend: nil,
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "go_framework", "backend"),
		},
		{
			name: "Unknown value in GoFramework option in Backend block",
			config: &config.Config{
				Backend:  &config.Backend{ModuleName: "project", GoFramework: "unknown", Port: 5000, Timeout: nil},
				Frontend: nil,
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigValueInOptionUnknown, "go_framework", "backend", "unknown"),
		},
		{
			name: "Port option in Backend block not found",
			config: &config.Config{
				Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 0, Timeout: nil},
				Frontend: nil,
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "port", "backend"),
		},
		{
			name: "Frontend block not found",
			config: &config.Config{
				Backend: &config.Backend{
					ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 5, Write: 10},
				},
				Frontend: nil,
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigBlockNotFound, "frontend"),
		},
		{
			name: "Package name in the Frontend block not found",
			config: &config.Config{
				Backend: &config.Backend{
					ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 5, Write: 10},
				},
				Frontend: &config.Frontend{
					PackageName: "", CSSFramework: "default", RuntimeEnvironment: "default",
					HTMX: "latest", Hyperscript: "latest",
				},
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "package_name", "frontend"),
		},
		{
			name: "CSS framework in the Frontend block not found",
			config: &config.Config{
				Backend: &config.Backend{
					ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 5, Write: 10},
				},
				Frontend: &config.Frontend{
					PackageName: "project", CSSFramework: "", RuntimeEnvironment: "default",
					HTMX: "latest", Hyperscript: "latest",
				},
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "css_framework", "frontend"),
		},
		{
			name: "CSS framework in the Frontend block has unknown value",
			config: &config.Config{
				Backend: &config.Backend{
					ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 5, Write: 10},
				},
				Frontend: &config.Frontend{
					PackageName: "project", CSSFramework: "unknown", RuntimeEnvironment: "default",
					HTMX: "latest", Hyperscript: "latest",
				},
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigValueInOptionUnknown, "css_framework", "frontend", "unknown"),
		},
		{
			name: "Runtime environment in the Frontend block not found",
			config: &config.Config{
				Backend: &config.Backend{
					ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 5, Write: 10},
				},
				Frontend: &config.Frontend{
					PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "",
					HTMX: "latest", Hyperscript: "latest",
				},
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "css_framework", "frontend"),
		},
		{
			name: "Runtime environment in the Frontend block has unknown value",
			config: &config.Config{
				Backend: &config.Backend{
					ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 5, Write: 10},
				},
				Frontend: &config.Frontend{
					PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "unknown",
					HTMX: "latest", Hyperscript: "latest",
				},
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigValueInOptionUnknown, "css_framework", "frontend", "unknown"),
		},
		{
			name: "HTMX in the Frontend block not found",
			config: &config.Config{
				Backend: &config.Backend{
					ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 5, Write: 10},
				},
				Frontend: &config.Frontend{
					PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "default",
					HTMX: "", Hyperscript: "latest",
				},
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "htmx", "frontend"),
		},
		{
			name: "Hyperscript in the Frontend block not found",
			config: &config.Config{
				Backend: &config.Backend{
					ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 5, Write: 10},
				},
				Frontend: &config.Frontend{
					PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "default",
					HTMX: "latest", Hyperscript: "",
				},
			},
			wantErr: fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "hyperscript", "frontend"),
		},
		{
			name: "Empty PackageName option in Frontend block",
			config: &config.Config{
				Backend: &config.Backend{
					ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 5, Write: 10},
				},
				Frontend: &config.Frontend{
					PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "default",
					HTMX: "latest", Hyperscript: "latest",
				},
			},
			wantErr: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr == "" {
				require.NoError(t, Validate(tc.config))
			} else {
				require.Error(t, Validate(tc.config), tc.wantErr)
			}
		})
	}
}
