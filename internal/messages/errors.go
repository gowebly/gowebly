package messages

const (
	/*
		List of the internal Gowebly CLI error messages.
	*/

	// ErrorGoweblyDINotComplete represents error message, when dependency injection not complete.
	ErrorGoweblyDINotComplete string = "dependency injection process is not complete: %w"

	/*
		List of the error messages for the current OS (folder, file, etc.).
	*/

	// ErrorOSFileIsExists represents error message, when user want to create a file, but this file is exists.
	ErrorOSFileIsExists string = "file '%s' is exists, can't be overwritten (clean up first)"
	// ErrorOSFolderIsExists represents error message, when user want to create a folder, but this folder is exists.
	ErrorOSFolderIsExists string = "folder with name '%s' is exists, can't be overwritten (clean up first)"
	// ErrorOSRemoveFile represents error message, when user want to remove a file.
	ErrorOSRemoveFile string = "can't remove a file with name '%s': %w"

	/*
		List of the error messages for cmd commands.
	*/

	// ErrorCMDNotExecuteCommand represents error message, when a cmd command was not executed.
	ErrorCMDNotExecuteCommand string = "can't execute command '%s' with option '%s': %w"

	/*
		List of the error messages for forms.
	*/

	// ErrorFormNotRun represents error message, when a form was not run.
	ErrorFormNotRun string = "failed to run '%s' form in the '%s' command: %w"

	/*
		List of the error messages for spinner.
	*/

	// ErrorSpinnerNotRun represents error message, when a spinner was not run.
	ErrorSpinnerNotRun string = "failed to run spinner for the '%s': %w"
)
