package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh/spinner"
	"github.com/gowebly/gowebly/v2/internal/actions"
	"github.com/gowebly/gowebly/v2/internal/forms"
	"github.com/gowebly/gowebly/v2/internal/helpers"
	"github.com/gowebly/gowebly/v2/internal/injectors"
	"github.com/gowebly/gowebly/v2/internal/messages"
	"github.com/gowebly/gowebly/v2/internal/variables"
)

// Create creates a new project with the given configuration.
func Create(di *injectors.Injector) error {
	// Add technical space for the 'create' command.
	fmt.Println()

	// Run create form.
	if err := forms.RunCreateForm(di); err != nil {
		return err
	}

	// Create a new context and a cancel function.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create buffered channel for one error value.
	errCh := make(chan error, 1)
	defer close(errCh)

	// Run action that creates project in a goroutine.
	go actions.CreateProjectAction(ctx, cancel, di, errCh)

	// Run spinner.
	if err := helpers.RunSpinnerWithContext(ctx, messages.CommandCreateSpinnerTitle, spinner.Line); err != nil {
		return fmt.Errorf(messages.ErrorSpinnerNotRun, "create", err)
	}

	// Handle potential error from action.
	if err := <-errCh; err != nil {
		return err
	}

	// Define variables.
	var goFramework, reactivityLibrary, cssFramework string

	// Check if the specified Go framework is valid.
	if _, ok := variables.ListGoFrameworks[di.Config.Backend.GoFramework]; ok {
		goFramework = variables.ListGoFrameworks[di.Config.Backend.GoFramework][1]
	}

	// Check if the specified frontend reactivity library is valid.
	if _, ok := variables.ListReactivityLibraries[di.Config.Frontend.ReactivityLibrary]; ok {
		reactivityLibrary = variables.ListReactivityLibraries[di.Config.Frontend.ReactivityLibrary][1]
	}

	// Check if the specified frontend CSS framework is valid.
	if _, ok := variables.ListCSSFrameworks[di.Config.Frontend.CSSFramework]; ok {
		cssFramework = variables.ListCSSFrameworks[di.Config.Frontend.CSSFramework][1]
	}

	// Generate content body.
	contentBody := fmt.Sprintf(
		messages.CommandCreateSummaryDescription,
		helpers.MakeStyled(messages.CommandCreateSummaryHeadingBackend, &helpers.StringStyle{Color: variables.ColorGrey}),
		helpers.MakeStyled(di.Config.Backend.ModuleName, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(goFramework, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(di.Config.Backend.Port, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(messages.CommandCreateSummaryHeadingFrontend, &helpers.StringStyle{Color: variables.ColorGrey}),
		helpers.MakeStyled(di.Config.Frontend.PackageName, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(reactivityLibrary, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(cssFramework, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(messages.CommandCreateSummaryHeadingTools, &helpers.StringStyle{Color: variables.ColorGrey}),
		helpers.MakeStyled(strconv.FormatBool(di.Config.Tools.IsUseAir), &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(strconv.FormatBool(di.Config.Tools.IsUseBun), &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(strconv.FormatBool(di.Config.Tools.IsUseTempl), &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(strconv.FormatBool(di.Config.Tools.IsUseGolangCILint), &helpers.StringStyle{Color: variables.ColorBlue}),
	)

	// Show created project info.
	fmt.Println(helpers.MakeStyled(
		messages.CommandCreateSummaryTitle,
		&helpers.StringStyle{Color: variables.ColorGreen, IsBold: true},
	))
	fmt.Println(helpers.MakeStyledFrame(
		contentBody,
		&helpers.FrameStyle{Padding: []int{1}, Color: variables.ColorGreen},
	))
	fmt.Println(helpers.MakeStyled(
		messages.CommandMoreInformationTitle,
		&helpers.StringStyle{Color: variables.ColorYellow},
	))

	return nil
}
