package commands

import "github.com/gowebly/gowebly/v2/internal/helpers"

// Run runs the 'run' cmd command.
func Run() error {
	return helpers.Execute(
		[]helpers.Command{
			{
				Name:    "air",
				Options: []string{"-c", ".air.toml"},
			},
		},
	)
}
