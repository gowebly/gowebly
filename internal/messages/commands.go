package messages

const (
	/*
		List of the spinner constants.
	*/

	// CommandCreateSpinnerTitle represents the title for the command create spinner.
	CommandCreateSpinnerTitle string = " Gowebly CLI is creating your project. Please wait..."

	/*
		List of the summary constants.
	*/

	// CommandUnknownSummaryTitle represents the title of the unknown summary.
	CommandErrorSummaryTitle string = "✕ Oops... Something went wrong!\n"

	// CommandCreateSummaryTitle represents the title of the project summary.
	CommandCreateSummaryTitle string = "✓ Your project has been created successfully!\n"
	// CommandCreateSummaryHeadingBackend represents the heading of the backend summary.
	CommandCreateSummaryHeadingBackend string = "Backend ↘"
	// CommandCreateSummaryHeadingFrontend represents the heading of the frontend summary.
	CommandCreateSummaryHeadingFrontend string = "Frontend ↘"
	// CommandCreateSummaryHeadingTools represents the heading of the tools summary.
	CommandCreateSummaryHeadingTools string = "Tools ↘"
	// CommandCreateSummaryDescription represents the description of the project summary.
	CommandCreateSummaryDescription string = "%s\n\nModule name in the go.mod file: %s\nGo web framework/router: %s\nServer port: %s\n\n%s\n\nPackage name in the package.json file: %s\nReactivity library: %s\nCSS framework: %s\n\n%s\n\nIs use Air tool to live-reloading? %s\nIs use Bun as a frontend runtime? %s\nIs use Templ to generate HTML? %s\nIs use golangci-lint to lint your Go code? %s"

	// CommandDoctorSummaryTitle represents the title of the system summary.
	CommandDoctorSummaryTitle string = "✓ Your system information has been collected successfully!\n"
	// CommandDoctorSummarySubTitle represents the subtitle of the system summary.
	CommandDoctorSummarySubTitle string = "Copy this information to paste into an issue ↘"
	// CommandDoctorSummaryDescription represents the description of the system summary.
	CommandDoctorSummaryDescription string = "%s\n\nGowebly CLI: %s (build with %s on %s/%s)\n\nGo: %s\nNode.js version: %s (npm %s)\nBun version: %s\n\nIs Air tool installed? %s\nIs Templ package installed? %s"

	// CommandUnknownSummaryTitle represents the title of the unknown summary.
	CommandUnknownSummaryTitle string = "~ This command doesn't exist...\n"
	// CommandUnknownSummarySubTitle represents the subtitle of the unknown summary.
	CommandUnknownSummarySubTitle string = "List of the available commands ↘"
	// CommandUnknownSummaryDescription represents the description of the unknown summary.
	CommandUnknownSummaryDescription string = "%s\n\n`%s` to start creating a new project\n`%s` to collecting info about your system\n`%s`    to running your project"

	/*
		List of the information constants.
	*/

	// CommandMoreInformationTitle represents the title of the more information string.
	CommandMoreInformationTitle string = "\n✱ For more information go to the official docs: https://gowebly.org \n"
)
