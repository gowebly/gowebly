package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gowebly/gowebly/v3/internal/actions"
	"github.com/gowebly/gowebly/v3/internal/forms"
	"github.com/gowebly/gowebly/v3/internal/helpers"
	"github.com/gowebly/gowebly/v3/internal/injectors"
	"github.com/gowebly/gowebly/v3/internal/messages"
	"github.com/gowebly/gowebly/v3/internal/variables"
)

// Create orchestrates the project creation workflow using the provided configuration.
func Create(di *injectors.Injector) error {
	fmt.Println()

	if err := forms.RunCreateForm(di); err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	defer close(errCh)

	go actions.CreateProjectAction(ctx, cancel, di, errCh)

	fmt.Println(helpers.MakeStyled(messages.CommandCreateWaitingTitle, &helpers.StringStyle{Color: variables.ColorYellow}))

	if err := <-errCh; err != nil {
		return err
	}

	var goFramework, reactivityLibrary, cssFramework string

	if _, ok := variables.ListGoFrameworks[di.Config.Backend.GoFramework]; ok {
		goFramework = variables.ListGoFrameworks[di.Config.Backend.GoFramework][1]
	}

	if _, ok := variables.ListReactivityLibraries[di.Config.Frontend.ReactivityLibrary]; ok {
		reactivityLibrary = variables.ListReactivityLibraries[di.Config.Frontend.ReactivityLibrary][1]
	}

	if _, ok := variables.ListCSSFrameworks[di.Config.Frontend.CSSFramework]; ok {
		cssFramework = variables.ListCSSFrameworks[di.Config.Frontend.CSSFramework][1]
	}

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
