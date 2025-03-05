package commands

import "github.com/gowebly/gowebly/v3/internal/helpers"

// Run runs the 'run' cmd command.
func Run() error {
	// Define running command by the 'Makefile'.
	runner := "make"
	runnerOptions := []string{"run"}

	// Check if Air tool config is exist.
	if helpers.IsExistInFolder(".air.toml", false) {
		// Define running command by the Air tool.
		runner = "air"
		runnerOptions = []string{"-c", ".air.toml"}
	}

	return helpers.Execute(
		[]helpers.Command{
			{
				Name:    runner,
				Options: runnerOptions,
			},
		},
	)
}
