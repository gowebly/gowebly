package constants

const (

	/*
		List of the common error messages.
	*/

	// ErrorDependencyInjectionNotComplete represents error message, when dependency injection
	// not complete.
	ErrorDependencyInjectionNotComplete string = "gowebly: a dependency injection process is not complete"
	// ErrorRunWithoutCommand represents error message, when user run cmd without
	// any command.
	ErrorRunWithoutCommand string = "gowebly: run without any commands or/and options"
	// ErrorRunUnknownCommand represents error message, when user run cmd with
	// unknown command.
	ErrorRunUnknownCommand string = "gowebly: run with unknown command"
	// ErrorProjectFolderIsNotEmpty represents error message, when user want to
	// create a new project, but the current folder is not empty.
	ErrorProjectFolderIsNotEmpty string = "project: folder is not empty, cannot be overwritten"

	/*
		List of the error messages for validators.
	*/

	// ErrorValidateConfigBackendBlockNotFound represents error message, when
	// 'backend' block in the config file is not found.
	ErrorValidateConfigBackendBlockNotFound string = "config: 'backend' block is required"
	// ErrorValidateConfigBackendNameNotFound represents error message, when
	// 'name' option in the 'backend' block in the config file is not found.
	ErrorValidateConfigBackendNameNotFound string = "config: 'name' option in the 'backend' block is required"
	// ErrorValidateConfigBackendPortNotFound represents error message, when
	// 'port' option in the 'backend' block in the config file is not found.
	ErrorValidateConfigBackendPortNotFound string = "config: 'port' option in the 'backend' block is required"
	// ErrorValidateConfigFrontendBlockNotFound represents error message, when
	// 'frontend' block in the config file is not found.
	ErrorValidateConfigFrontendBlockNotFound string = "config: 'frontend' block is required"
	// ErrorValidateConfigFrontendHTMXNotFound represents error message, when
	// 'htmx' option in the 'frontend' block in the config file is not found.
	ErrorValidateConfigFrontendHTMXNotFound string = "config: 'htmx' option in the 'frontend' block is required"
	// ErrorValidateConfigFrontendHyperscriptNotFound represents error message, when
	// 'hyperscript' option in the 'frontend' block in the config file is not found.
	ErrorValidateConfigFrontendHyperscriptNotFound string = "config: 'hyperscript' option in the 'frontend' block is required"
)
