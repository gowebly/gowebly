package constants

const (

	/*
		List of the common error messages.
	*/

	// ErrorDependencyInjectionNotComplete represents error message, when dependency injection
	// not complete.
	ErrorDependencyInjectionNotComplete string = "gowebly: a dependency injection process is not complete"
	// ErrorRunWithoutFlags represents error message, when user run cmd without
	// flags.
	ErrorRunWithoutFlags string = "gowebly: run without any commands and options"
	// ErrorRunUnknownCommand represents error message, when user run cmd with
	// unknown command.
	ErrorRunUnknownCommand string = "gowebly: run with unknown command"

	/*
		List of the error messages for 'add' command.
	*/

	// ErrorRunAddCommandWithoutCSSFrameworkName represents error message, when user run
	// add command without CSS framework name.
	ErrorRunAddCommandWithoutCSSFrameworkName string = "'add' command: run without CSS framework name option"
	// ErrorRunAddCommandWithUnknownCSSFrameworkName represents error message, when user run
	// add command with an unknown CSS framework name.
	ErrorRunAddCommandWithUnknownCSSFrameworkName string = "'add' command: run with unknown CSS framework name"
)
