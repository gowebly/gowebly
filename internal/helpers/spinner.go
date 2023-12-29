package helpers

import (
	"context"

	"github.com/charmbracelet/huh/spinner"
)

// RunSpinnerWithContext runs the spinner with the given context.
func RunSpinnerWithContext(ctx context.Context, title string, spinnerType spinner.Type) error {
	return spinner.New().Type(spinnerType).Title(title).Context(ctx).Run()
}
