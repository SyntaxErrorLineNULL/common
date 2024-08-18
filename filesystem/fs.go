package filesystem

import (
	"os"
	"path/filepath"
)

// RecursiveCreatePath ensures that all directories in the specified file path exist.
// If any directories in the path do not exist, it recursively creates them.
func RecursiveCreatePath(filePath string) error {
	// Extract the directory part of the file path.
	dirname := filepath.Dir(filePath)

	// Check if the directory exists.
	// If it does not exist, `os.Stat` returns an error which we check using `os.IsNotExist`.
	if _, err := os.Stat(dirname); !os.IsNotExist(err) {
		// If the directory exists or some other error occurred (not `os.IsNotExist`), return the error.
		return err
	}
	// Recursively call `RecursiveCreatePath` to create parent directories.
	// This ensures that the entire directory path leading up to `dirname` is created.
	if err := RecursiveCreatePath(dirname); err != nil {
		// If an error occurs while creating parent directories, return the error.
		return err
	}
	// Create the directory with permissions set to 0755 (read/write/execute for owner, read/execute for others).
	if err := os.Mkdir(dirname, 0o755); err != nil {
		// If an error occurs while creating the directory, return the error.
		return err
	}

	// Return nil to indicate success.
	return nil
}
