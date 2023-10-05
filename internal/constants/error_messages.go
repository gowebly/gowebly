package constants

const (
	/*
		List of the common error messages.
	*/

	// ErrorDependencyInjectionNotComplete represents error message, when dependency injection not complete.
	ErrorDependencyInjectionNotComplete string = "gowebly: a dependency injection process is not complete"
	// ErrorRunCommandWithUnknownFlag represents error message, when user run cmd with unknown flag.
	ErrorRunCommandWithUnknownFlag string = "gowebly: unknown flag '%s' of the '%s' command"
	// ErrorProjectFolderIsNotEmpty represents error message, when user want to
	// create a new project, but the current folder is not empty.
	ErrorProjectFolderIsNotEmpty string = "os: a current folder is not empty, cannot be overwritten"

	/*
		List of the error messages for validators.
	*/

	// ErrorValidateConfigBackendBlockNotFound represents error message, when a block in the config file is not found.
	ErrorValidateConfigBlockNotFound string = "config: '%s' block is required"
	// ErrorValidateConfigOptionInBlockNotFound represents error message, when an option in the block is not found.
	ErrorValidateConfigOptionInBlockNotFound string = "config: '%s' option in the '%s' block is required"
	// ErrorValidateConfigValueInOptionUnknown represents error message, when an option in the block has unknown value.
	ErrorValidateConfigValueInOptionUnknown string = "config: '%s' option in the '%s' block has unknown value '%s'"
)
