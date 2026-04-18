package variables

// Available options for the interactive project creation form.
// Each map key is the option ID, value is [option_value, display_label].
var (
	ListGoFrameworks map[string][]string = map[string][]string{
		"default":    {"default", "Built-in net/http package"},
		"fiber":      {"fiber", "Fiber"},
		"gin":        {"gin", "Gin"},
		"echo":       {"echo", "Echo"},
		"chi":        {"chi", "Chi"},
		"httprouter": {"httprouter", "HttpRouter"},
		"gorilla":    {"gorilla", "Gorilla/Mux"},
		"pocketbase": {"pocketbase", "PocketBase"},
	}

	ListReactivityLibraries map[string][]string = map[string][]string{
		"htmx":             {"htmx", "htmx"},
		"htmx_hyperscript": {"htmx_hyperscript", "htmx with hyperscript"},
		"htmx_alpinejs":    {"htmx_alpinejs", "htmx with Alpine.js"},
	}

	ListCSSFrameworks map[string][]string = map[string][]string{
		"default":     {"default", "Default styles"},
		"tailwindcss": {"tailwindcss", "Tailwind CSS"},
		"daisyui":     {"daisyui", "Tailwind CSS with daisyUI components"},
		"flowbite":    {"flowbite", "Tailwind CSS with Flowbite components"},
		"prelineui":   {"prelineui", "Tailwind CSS with Preline UI components"},
		"unocss":      {"unocss", "UnoCSS"},
		"bootstrap":   {"bootstrap", "Bootstrap"},
		"bulma":       {"bulma", "Bulma"},
	}
)
