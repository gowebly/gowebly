package commands

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/gowebly/gowebly/v3/internal/helpers"
	"github.com/gowebly/gowebly/v3/internal/messages"
	"github.com/gowebly/gowebly/v3/internal/variables"
)

// Doctor displays system information and checks for installed dependencies.
func Doctor() error {
	fmt.Println()

	goVersion, goInstalled := helpers.GetToolVersionSafe("go", "version")
	if !goInstalled {
		goVersion = "not installed"
	}

	nodeVersion, nodeInstalled := helpers.GetToolVersionSafe("node", "-v")
	if !nodeInstalled {
		nodeVersion = "not installed"
	}

	npmVersion, npmInstalled := helpers.GetToolVersionSafe("npm", "-v")
	if !npmInstalled {
		npmVersion = "not installed"
	}

	bunVersion, bunInstalled := helpers.GetToolVersionSafe("bun", "-v")
	if !bunInstalled {
		bunVersion = "not installed"
	}

	_, airInstalled := helpers.GetToolVersionSafe("air", "-v")

	_, templInstalled := helpers.GetToolVersionSafe("templ", "version")

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
		helpers.MakeStyled(strconv.FormatBool(airInstalled), &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled(strconv.FormatBool(templInstalled), &helpers.StringStyle{Color: variables.ColorBlue}),
	)

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
