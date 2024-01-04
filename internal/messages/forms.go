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
	// FormGoFrameworkDefaultKey represents the key for the built-in net/http package select.
	FormGoFrameworkDefaultKey string = "Built-in net/http"
	// FormGoFrameworkDefaultKey represents the value for the built-in net/http package select.
	FormGoFrameworkDefaultValue string = "default"
	// FormGoFrameworkFiberKey represents the key for the Fiber framework select.
	FormGoFrameworkFiberKey string = "Fiber"
	// FormGoFrameworkFiberValue represents the value for the Fiber framework select.
	FormGoFrameworkFiberValue string = "fiber"
	// FormGoFrameworkGinKey represents the key for the Gin framework select.
	FormGoFrameworkGinKey string = "Gin"
	// FormGoFrameworkGinValue represents the value for the Gin framework select.
	FormGoFrameworkGinValue string = "gin"
	// FormGoFrameworkEchoKey represents the key for the Echo framework select.
	FormGoFrameworkEchoKey string = "Echo"
	// FormGoFrameworkEchoValue represents the value for the Echo framework select.
	FormGoFrameworkEchoValue string = "echo"
	// FormGoFrameworkHttpRouterKey represents the key for the HttpRouter framework select.
	FormGoFrameworkHttpRouterKey string = "HttpRouter"
	// FormGoFrameworkHttpRouterValue represents the value for the HttpRouter framework select.
	FormGoFrameworkHttpRouterValue string = "httprouter"

	// FormReactiveLibraryTitle represents the title for the frontend reactive library select.
	FormReactiveLibraryTitle string = "Select a reactivity framework or library\n"
	// FormReactiveLibraryDescription represents the description for the frontend reactive library select.
	FormReactiveLibraryDescription string = "This framework (or library) will be used to build the frontend part\nof your application.\n"
	// FormReactiveLibraryVanillaKey represents the key for the pure (vanilla) TypeScript select.
	FormReactiveLibraryVanillaKey string = "Pure (vanilla) TypeScript"
	// FormReactiveLibraryVanillaValue represents the value for the pure (vanilla) TypeScript select.
	FormReactiveLibraryVanillaValue string = "vanilla"
	// FormReactiveLibraryVueKey represents the key for the Vue select.
	FormReactiveLibraryVueKey string = "Vue"
	// FormReactiveLibraryVueValue represents the value for the Vue select.
	FormReactiveLibraryVueValue string = "vue"
	// FormReactiveLibraryNuxtKey represents the key for the Nuxt select.
	FormReactiveLibraryNuxtKey string = "Nuxt"
	// FormReactiveLibraryNuxtValue represents the value for the Nuxt select.
	FormReactiveLibraryNuxtValue string = "nuxt"
	// FormReactiveLibraryReactKey represents the key for the React select.
	FormReactiveLibraryReactKey string = "React"
	// FormReactiveLibraryReactValue represents the value for the React select.
	FormReactiveLibraryReactValue string = "react"
	// FormReactiveLibraryNextKey represents the key for the Next select.
	FormReactiveLibraryNextKey string = "Next"
	// FormReactiveLibraryNextValue represents the value for the Next select.
	FormReactiveLibraryNextValue string = "next"
	// FormReactiveLibrarySvelteKey represents the key for the Svelte select.
	FormReactiveLibrarySvelteKey string = "Svelte"
	// FormReactiveLibrarySvelteValue represents the value for the Svelte select.
	FormReactiveLibrarySvelteValue string = "svelte"
	// FormReactiveLibrarySvelteKitKey represents the key for the SvelteKit select.
	FormReactiveLibrarySvelteKitKey string = "SvelteKit"
	// FormReactiveLibrarySvelteKitValue represents the value for the SvelteKit select.
	FormReactiveLibrarySvelteKitValue string = "sveltekit"

	// FormCSSFrameworkTitle represents the title for the frontend CSS framework select.
	FormCSSFrameworkTitle string = "Select a CSS framework or components library\n"
	// FormCSSFrameworkDescription represents the description for the frontend CSS framework select.
	FormCSSFrameworkDescription string = "This framework (or components library) will be used to build the UI part\nof your application.\n"
	// FormCSSFrameworkDefaultKey represents the key for the default styles select.
	FormCSSFrameworkDefaultKey string = "Default styles"
	// FormCSSFrameworkDefaultValue represents the value for the default styles select.
	FormCSSFrameworkDefaultValue string = "default"
	// FormCSSFrameworkTailwindCSSKey represents the key for the Tailwind CSS framework select.
	FormCSSFrameworkTailwindCSSKey string = "Tailwind CSS"
	// FormCSSFrameworkTailwindCSSValue represents the value for the Tailwind CSS framework select.
	FormCSSFrameworkTailwindCSSValue string = "tailwindcss"
	// FormCSSFrameworkDaisyUIKey represents the key for the Tailwind CSS with daisyUI components select.
	FormCSSFrameworkDaisyUIKey string = "Tailwind CSS with daisyUI components"
	// FormCSSFrameworkDaisyUIValue represents the value for the Tailwind CSS with daisyUI components select.
	FormCSSFrameworkDaisyUIValue string = "daisyui"
	// FormCSSFrameworkFlowbiteKey represents the key for the Tailwind CSS with Flowbite components select.
	FormCSSFrameworkFlowbiteKey string = "Tailwind CSS with Flowbite components"
	// FormCSSFrameworkFlowbiteValue represents the value for the Tailwind CSS with Flowbite components select.
	FormCSSFrameworkFlowbiteValue string = "flowbite"
	// FormCSSFrameworkUnoCSSKey represents the key for the UnoCSS select.
	FormCSSFrameworkUnoCSSKey string = "UnoCSS"
	// FormCSSFrameworkUnoCSSValue represents the value for the UnoCSS select.
	FormCSSFrameworkUnoCSSValue string = "unocss"
	// FormCSSFrameworkBootstrapKey represents the key for the Bootstrap framework select.
	FormCSSFrameworkBootstrapKey string = "Bootstrap"
	// FormCSSFrameworkBootstrapValue represents the value for the Bootstrap framework select.
	FormCSSFrameworkBootstrapValue string = "bootstrap"
	// FormCSSFrameworkBulmaKey represents the key for the Bulma framework select.
	FormCSSFrameworkBulmaKey string = "Bulma"
	// FormCSSFrameworkBulmaValue represents the value for the Bulma framework select.
	FormCSSFrameworkBulmaValue string = "bulma"

	/*
		Form elements: confirm fields.
	*/

	// FormAirUsageTitle represents the title for the Air tool switch.
	FormAirUsageTitle string = "Use the Air tool to enable live-reloading for your application?"

	// FormTempleUsageTitle represents the title for the Templ package switch.
	FormTemplUsageTitle string = "Use the Templ package to build HTML templates with Go?"

	// FormHyperscriptUsageTitle represents the title for the hyperscript switch.
	FormHyperscriptUsageTitle string = "Use the Hyperscript library on the frontend?"

	// FormRuntimeUsageTitle represents the title for the Bun runtime environment switch.
	FormBunUsageTitle string = "Use the Bun as a frontend runtime environment?"
)
