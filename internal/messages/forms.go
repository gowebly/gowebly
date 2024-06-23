package messages

const (
	/*
		Top-level form constants.
	*/

	// FormPromptSignature represents the prompt signature for the form.
	FormPromptSignature string = "> "

	/*
		Form elements: note fields.
	*/

	// FormWelcomeDescription represents the description of the welcome message.
	FormWelcomeDescription string = "A next-generation CLI tool that makes it easy to create\namazing web applications with Go on the backend, using\nhtmx, hyperscript or Alpine.js, and the most popular\nCSS frameworks on the frontend."

	/*
		Form elements: input fields.
	*/

	// FormGoModuleNameTitle represents the title for the Go module name input.
	FormGoModuleNameTitle string = "What's your Go module name in the go.mod file?\n"
	// FormGoModuleNameDescription represents the description for the Go module name input.
	FormGoModuleNameDescription string = "Option can be any name of your Go module (e.g., `github.com/user/project`).\n"

	// FormPortTitle represents the title for the backend port input.
	FormPortTitle string = "What's your port number to run backend server?\n"
	// FormPortDescription represents the description for the backend port input.
	FormPortDescription string = "Option can be any available port in your system (e.g., `7000`).\n"

	// FormPackageNameTitle represents the title for the frontend package name input.
	FormPackageNameTitle string = "What's your project name in the package.json file?\n"
	// FormPackageNameDescription represents the description for the frontend package name input.
	FormPackageNameDescription string = "Option can be any name of your package.json (e.g., `project`).\n"

	/*
		Form elements: select fields.
	*/

	// FormGoFrameworkTitle represents the title for the Go framework select.
	FormGoFrameworkTitle string = "Select the Go web framework or router\n"
	// FormGoFrameworkDescription represents the description for the Go framework select.
	FormGoFrameworkDescription string = "This framework (or router) will be used to build\nthe backend part of your application.\n"

	// FormGoFrameworkTitle represents the title for the reactivity library select.
	FormReactivityLibraryTitle string = "Select the reactivity library\n"
	// FormGoFrameworkDescription represents the description for the reactivity library select.
	FormReactivityLibraryDescription string = "This reactivity library will be used in the frontend part\nof your application.\n"

	// FormCSSFrameworkTitle represents the title for the frontend CSS framework select.
	FormCSSFrameworkTitle string = "Select a CSS framework or components library\n"
	// FormCSSFrameworkDescription represents the description for the frontend CSS framework select.
	FormCSSFrameworkDescription string = "This framework (or components library) will be used to build\nthe UI part of your application.\n"

	/*
		Form elements: confirm fields.
	*/

	// FormAirUsageTitle represents the title for the Air tool switch.
	FormAirUsageTitle string = "Use the Air tool to enable live-reloading for your application?\n"
	// FormAirUsageDescription represents the description for the Air tool switch.
	FormAirUsageDescription string = "This tool will enable live-reloading mode for your backend\nand frontend code of your application.\n\nFor more info → https://github.com/cosmtrek/air"

	// FormRuntimeUsageTitle represents the title for the Bun runtime environment switch.
	FormBunUsageTitle string = "Use the Bun as a frontend runtime environment?\n"
	// FormRuntimeUsageDescription represents the description for the Bun runtime environment switch.
	FormBunUsageDescription string = "This environment will be used to build the frontend\npart of your application.\n\nFor more info → https://github.com/oven-sh/bun"

	// FormTempleUsageTitle represents the title for the Templ package switch.
	FormTemplUsageTitle string = "Use the Templ package to build HTML templates with Go?\n"
	// FormTempleUsageDescription represents the description for the Templ package switch.
	FormTemplUsageDescription string = "This package will be used to generate HTML templates\nby Go for your application.\n\nFor more info → https://github.com/a-h/templ"

	// FormGolangCILintUsageTitle represents the title for the Golang CI Lint switch.
	FormGolangCILintUsageTitle string = "Use the Golang CI Lint to lint your Go code?\n"
	// FormGolangCILintUsageDescription represents the description for the Golang CI Lint switch.
	FormGolangCILintUsageDescription string = "This tool will be used to lint your Go code.\n\nFor more info → https://github.com/golangci/golangci-lint"
)
