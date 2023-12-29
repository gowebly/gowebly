package messages

const (
	/*
		Top-level Form constants.
	*/

	// FormPromptSignature represents the prompt signature for the form.
	FormPromptSignature string = "> "

	/*
		Form elements: note fields.
	*/

	// FormWelcomeDescription represents the description of the welcome message.
	FormWelcomeDescription string = "A next-generation CLI tool that makes it easy to create amazing web applications\nwith Go on the backend, using htmx & hyperscript or the most popular TypeScript\nand CSS frameworks on the frontend."

	/*
		Form elements: input fields.
	*/

	// FormGoModuleNameTitle represents the title for the Go module name input.
	FormGoModuleNameTitle string = "What's your Go module name in the go.mod file?\n"
	// FormGoModuleNameDescription represents the description for the Go module name input.
	FormGoModuleNameDescription string = "Option can be any name of your Go module, for example, `github.com/user/project`.\n"

	// FormPortTitle represents the title for the backend port input.
	FormPortTitle string = "What's your port number to run backend server?\n"
	// FormPortDescription represents the description for the backend port input.
	FormPortDescription string = "Option can be any available port in your system, for example, `7000`.\n"

	// FormPackageNameTitle represents the title for the frontend package name input.
	FormPackageNameTitle string = "What's your project name in the package.json file?\n"
	// FormPackageNameDescription represents the description for the frontend package name input.
	FormPackageNameDescription string = "Option can be any name of your package.json, for example, `project`.\n"

	/*
		Form elements: select fields.
	*/

	// FormGoFrameworkTitle represents the title for the Go framework select.
	FormGoFrameworkTitle string = "Select the Go web framework or router\n"
	// FormGoFrameworkDescription represents the description for the Go framework select.
	FormGoFrameworkDescription string = "This framework (or router) will be used to build the backend part\nof your application.\n"
	// FormGoFrameworkFiberTitle represents the title for the built-in net/http package select.
	FormGoFrameworkDefaultTitle string = "Built-in net/http"
	// FormGoFrameworkFiberTitle represents the title for the Fiber framework select.
	FormGoFrameworkFiberTitle string = "Fiber"
	// FormGoFrameworkGinTitle represents the title for the Gin framework select.
	FormGoFrameworkGinTitle string = "Gin"
	// FormGoFrameworkEchoTitle represents the title for the Echo framework select.
	FormGoFrameworkEchoTitle string = "Echo"
	// FormGoFrameworkHttpRouterTitle represents the title for the HttpRouter framework select.
	FormGoFrameworkHttpRouterTitle string = "HttpRouter"

	// FormReactiveLibraryTitle represents the title for the frontend reactive library select.
	FormReactiveLibraryTitle string = "Select a reactivity framework or library\n"
	// FormReactiveLibraryDescription represents the description for the frontend reactive library select.
	FormReactiveLibraryDescription string = "This framework (or library) will be used to build the frontend part\nof your application.\n"
	// FormReactiveLibraryVanillaTitle represents the title for the pure (vanilla) TypeScript select.
	FormReactiveLibraryVanillaTitle string = "Pure (vanilla) TypeScript"
	// FormReactiveLibraryVueTitle represents the title for the Vue select.
	FormReactiveLibraryVueTitle string = "Vue"
	// FormReactiveLibraryNuxtTitle represents the title for the Nuxt select.
	FormReactiveLibraryNuxtTitle string = "Nuxt"
	// FormReactiveLibraryReactTitle represents the title for the React select.
	FormReactiveLibraryReactTitle string = "React"
	// FormReactiveLibraryNextTitle represents the title for the Next select.
	FormReactiveLibraryNextTitle string = "Next"
	// FormReactiveLibrarySvelteTitle represents the title for the Svelte select.
	FormReactiveLibrarySvelteTitle string = "Svelte"
	// FormReactiveLibrarySvelteKitTitle represents the title for the SvelteKit select.
	FormReactiveLibrarySvelteKitTitle string = "SvelteKit"

	// FormCSSFrameworkTitle represents the title for the frontend CSS framework select.
	FormCSSFrameworkTitle string = "Select a CSS framework or components library\n"
	// FormCSSFrameworkDescription represents the description for the frontend CSS framework select.
	FormCSSFrameworkDescription string = "This framework (or components library) will be used to build the UI part\nof your application.\n"
	// FormCSSFrameworkTailwindTitle represents the title for the default styles select.
	FormCSSFrameworkDefaultTitle string = "Default styles"
	// FormCSSFrameworkTailwindCSSTitle represents the title for the Tailwind CSS framework select.
	FormCSSFrameworkTailwindCSSTitle string = "Tailwind CSS"
	// FormCSSFrameworkDaisyUITitle represents the title for the Tailwind CSS with daisyUI components select.
	FormCSSFrameworkDaisyUITitle string = "Tailwind CSS with daisyUI components"
	// FormCSSFrameworkFlowbiteTitle represents the title for the Tailwind CSS with Flowbite components select.
	FormCSSFrameworkFlowbiteTitle string = "Tailwind CSS with Flowbite components"
	// FormCSSFrameworkBootstrapTitle represents the title for the Bootstrap framework select.
	FormCSSFrameworkBootstrapTitle string = "Bootstrap"
	// FormCSSFrameworkBulmaTitle represents the title for the Bulma framework select.
	FormCSSFrameworkBulmaTitle string = "Bulma"

	/*
		Form elements: confirm fields.
	*/

	// FormAirUsageTitle represents the title for the Air tool switch.
	FormAirUsageTitle string = "Use the Air tool to enable hot-reloading in the application?"

	// FormTempleUsageTitle represents the title for the Templ package switch.
	FormTemplUsageTitle string = "Use the Templ package to build HTML templates with Go?"

	// FormHTMXUsageTitle represents the title for the htmx as reactive library switch.
	FormHTMXUsageTitle string = "Use the htmx as a reactive frontend library?"

	// FormHyperscriptUsageTitle represents the title for the hyperscript library switch.
	FormHyperscriptUsageTitle string = "Add the hyperscript library to the frontend?"

	// FormRuntimeUsageTitle represents the title for the Bun JavaScript runtime environment switch.
	FormBunUsageTitle string = "Use the Bun as a JavaScript/TypeScript runtime environment?"
)
