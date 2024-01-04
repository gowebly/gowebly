package commands

import "github.com/gowebly/gowebly/internal/helpers"

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
