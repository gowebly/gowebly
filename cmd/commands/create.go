package commands

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/huh/spinner"
	"github.com/gowebly/gowebly/internal/actions"
	"github.com/gowebly/gowebly/internal/forms"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
	"github.com/gowebly/gowebly/internal/messages"
	"github.com/gowebly/gowebly/internal/variables"
)

// Create creates a new project with the given configuration.
func Create(di *injectors.Injector) error {
	// Add technical space for 'create' co.
	fmt.Println()

	// Run welcome note form.
	if err := forms.CreateWelcomeForm(di); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "welcome", "create", err)
	}

	// Run project settings form.
	if err := forms.CreateProjectSettingsForm(di); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "project settings", "create", err)
	}

	// Run Go framework form.
	if err := forms.CreateGoFrameworkForm(di); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "go framework", "create", err)
	}

	// Run HTMX form.
	if err := forms.CreateHTMXForm(di); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "htmx", "create", err)
	}

	// Check, if HTMX is used.
	if di.Config.Frontend.IsUseHTMX {
		// If yes, run Templ form.
		if err := forms.CreateTemplForm(di); err != nil {
			return fmt.Errorf(messages.ErrorFormNotRun, "templ", "create", err)
		}
	} else {
		// If not, run reactive library form.
		if err := forms.CreateReactiveLibraryForm(di); err != nil {
			return fmt.Errorf(messages.ErrorFormNotRun, "reactive library", "create", err)
		}
	}

	// Run CSS framework form.
	if err := forms.CreateCSSFrameworkForm(di); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "css framework", "create", err)
	}

	// TODO: add function to create new project.
	action := func() error {
		// Create a new folder(s).
		if err := helpers.MakeFolders("assets", "static", "templates/pages"); err != nil {
			return err
		}

		// Create backend, deploy and misc files from templates.
		if err := actions.CreateBackendFiles(di); err != nil {
			return err
		}

		// Create frontend files.
		if err := actions.CreateFrontendFiles(di); err != nil {
			return err
		}

		return nil
	}

	// Create a new context and a cancel function.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a new action to create a new project.
	go action()

	// Show spinner.
	if err := helpers.RunSpinnerWithContext(ctx, messages.CommandCreateSpinnerTitle, spinner.Line); err != nil {
		return fmt.Errorf(messages.ErrorSpinnerNotRun, "create", err)
	}

	// Generate content body.
	contentBody := fmt.Sprintf(
		messages.CommandCreateSummaryDescription,
		helpers.MakeStyled(messages.CommandCreateSummaryHeadingBackend, &helpers.StringStyle{Color: variables.ColorGrey}),
		helpers.MakeStyled(di.Config.Backend.GoFramework, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(di.Config.Backend.Port, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(messages.CommandCreateSummaryHeadingFrontend, &helpers.StringStyle{Color: variables.ColorGrey}),
		helpers.MakeStyled(di.Config.Frontend.ReactiveLibrary, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(di.Config.Frontend.CSSFramework, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(messages.CommandCreateSummaryHeadingTools, &helpers.StringStyle{Color: variables.ColorGrey}),
		helpers.MakeStyled(strconv.FormatBool(di.Config.Tools.IsUseAir), &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(strconv.FormatBool(di.Config.Tools.IsUseBun), &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(strconv.FormatBool(di.Config.Tools.IsUseTempl), &helpers.StringStyle{Color: variables.ColorBlue}),
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
