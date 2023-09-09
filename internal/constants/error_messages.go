package constants

const (
	/*
		List of the common error messages.
	*/

	// ErrorRunWithoutFlags represents error message, when user run cmd without
	// flags.
	ErrorRunWithoutFlags string = ""
	// ErrorRunUnknownCommand represents error message, when user run cmd with
	// unknown command.
	ErrorRunUnknownCommand string = ""

	/*
		List of the error messages for 'add' command.
	*/

	// ErrorRunAddCommandWithoutCSSFrameworkName represents error message, when user run
	// add command without CSS framework name.
	ErrorRunAddCommandWithoutCSSFrameworkName string = ""
	// ErrorRunAddCommandWithUnknownCSSFrameworkName represents error message, when user run
	// add command with an unknown CSS framework name.
	ErrorRunAddCommandWithUnknownCSSFrameworkName string = ""
)
