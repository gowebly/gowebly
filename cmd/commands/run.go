package commands

import "github.com/gowebly/gowebly/v3/internal/helpers"

// Run starts the development server using Air or Make.
func Run() error {
	runner := "make"
	runnerOptions := []string{"run"}

	if helpers.IsExistInFolder(".air.toml", false) {
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
