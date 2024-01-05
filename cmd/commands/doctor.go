package commands

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/gowebly/gowebly/v2/internal/helpers"
	"github.com/gowebly/gowebly/v2/internal/messages"
	"github.com/gowebly/gowebly/v2/internal/variables"
)

// Doctor runs the 'doctor' cmd command.
func Doctor() error {
	// Add technical space for 'doctor' command.
	fmt.Println()

	// Get Go version.
	goVersion, _ := helpers.GetToolVersion("go", "version")
	if goVersion == "" {
		goVersion = "not installed"
	}

	// Get Node.js version.
	nodeVersion, _ := helpers.GetToolVersion("node", "-v")
	if nodeVersion == "" {
		nodeVersion = "not installed"
	}

	// Get npm version.
	npmVersion, _ := helpers.GetToolVersion("npm", "-v")
	if npmVersion == "" {
		npmVersion = "not installed"
	}

	// Get Bun version.
	bunVersion, _ := helpers.GetToolVersion("bun", "-v")
	if bunVersion == "" {
		bunVersion = "not installed"
	}

	// Get Air version.
	airVersion, _ := helpers.CheckToolIsInstalled("air", "-v")

	// Get Templ version.
	templVersion, _ := helpers.CheckToolIsInstalled("templ", "version")

	// Generate content body.
	contentBody := fmt.Sprintf(
		messages.CommandDoctorSummaryDescription,
		helpers.MakeStyled(messages.CommandDoctorSummarySubTitle, &helpers.StringStyle{Color: variables.ColorGrey}),
		helpers.MakeStyled(variables.GoweblyVersion, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(runtime.Version(), &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(runtime.GOOS, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(runtime.GOARCH, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(strings.ReplaceAll(goVersion, "go version go", ""), &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(strings.ReplaceAll(nodeVersion, "v", ""), &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(npmVersion, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(bunVersion, &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(strconv.FormatBool(airVersion), &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(strconv.FormatBool(templVersion), &helpers.StringStyle{Color: variables.ColorBlue}),
	)

	// Show created project info.
	fmt.Println(helpers.MakeStyled(
		messages.CommandDoctorSummaryTitle,
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
