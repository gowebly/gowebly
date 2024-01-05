package helpers

import (
	"context"
	"testing"

	"github.com/charmbracelet/huh/spinner"
)

func TestRunSpinnerWithContext(t *testing.T) {
	ctx := context.TODO()

	// Run spinner.
	if err := RunSpinnerWithContext(ctx, "title", spinner.Line); err != nil {
		t.Errorf("RunSpinnerWithContext() error = %v", err)
	}
}
