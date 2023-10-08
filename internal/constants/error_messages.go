package constants

const (
	/*
		List of the internal gowebly error messages.
	*/

	// ErrorGoweblyConfigPathIsEmpty represents error message, when a config path is empty.
	ErrorGoweblyConfigPathIsEmpty string = "gowebly: a path of the config file is empty"
	// ErrorGoweblyDINotComplete represents error message, when dependency injection not complete.
	ErrorGoweblyDINotComplete string = "gowebly: a dependency injection process is not complete"
	// ErrorGoweblyUnmarshalConfigFileToStructNotComplete represents error message, when unmarshal process is not complete.
	ErrorGoweblyUnmarshalConfigFileToStructNotComplete string = "gowebly: unmarshal data from the YAML config file to struct is not complete"
	// ErrorGoweblyDefaultYAMLConfigFileNotFound represents error message, when a default config file is not found.
	ErrorGoweblyDefaultYAMLConfigFileNotFound string = "gowebly: a default YAML config file is not found"
	// ErrorGoweblyYAMLConfigFileNotValid represents error message, when a structure of the config file is not valid.
	ErrorGoweblyYAMLConfigFileNotValid string = "gowebly: a structure of the YAML config file is not valid"

	/*
		List of the error messages for the current OS (folder, file, etc.)..
	*/

	// ErrorOSFileIsExists represents error message, when user want to create a file, but this file is exists.
	ErrorOSFileIsExists string = "os: file '%s' is exists, can't be overwritten (clean up first)"
	// ErrorOSFolderIsExists represents error message, when user want to create a folder, but this folder is exists.
	ErrorOSFolderIsExists string = "os: folder with name '%s' is exists, can't be overwritten (clean up first)"
	// ErrorOSRemoveFile represents error message, when user want to remove a file.
	ErrorOSRemoveFile string = "os: can't remove a file with name '%s'"

	/*
		List of the error messages for HTTP client.
	*/

	// ErrorHTTPDownloadFile represents error message, when user want to download a file from URL.
	ErrorHTTPDownloadFile string = "http: can't download file from URL '%s' (code %d)"

	/*
		List of the error messages for cmd commands.
	*/

	// ErrorCMDUnknownFlag represents error message, when user run cmd command with unknown flag.
	ErrorCMDUnknownFlag string = "cmd: unknown flag '%s' for the '%s' command"
	// ErrorCMDNotExecuteCommand represents error message, when a cmd command was not executed.
	ErrorCMDNotExecuteCommand string = "cmd: can't execute command '%s' with options %s"

	/*
		List of the error messages for validators.
	*/

	// ErrorValidateConfigBlockNotFound represents error message, when a block in the config file is not found.
	ErrorValidateConfigBlockNotFound string = "config: '%s' block is required"
	// ErrorValidateConfigOptionInBlockNotFound represents error message, when an option in the block is not found.
	ErrorValidateConfigOptionInBlockNotFound string = "config: '%s' option in the '%s' block is required"
	// ErrorValidateConfigValueInOptionUnknown represents error message, when an option in the block has unknown value.
	ErrorValidateConfigValueInOptionUnknown string = "config: '%s' option in the '%s' block has unknown value '%s'"

	/*
		List of the error messages for helpers.
	*/

	// ErrorHelperToolNotInstalled represents error message, when a tool in the current system is not installed.
	ErrorHelperToolNotInstalled string = "'%s' tool is not installed on your system"
)
