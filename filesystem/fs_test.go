package filesystem

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRecursiveCreatePath tests the RecursiveCreatePath function to ensure it correctly creates directories for a given file path.
func TestRecursiveCreatePath(t *testing.T) {
	t.Parallel()

	// CreateNestedDirectories tests the behavior of the RecursiveCreatePath function
	// when creating a deeply nested directory structure. The test ensures that the function
	// correctly creates all the specified directories and handles directory creation without errors.
	t.Run("CreateNestedDirectories", func(t *testing.T) {
		// Define the base directory where the nested directories will be created.
		// t.TempDir() creates a temporary directory for the test, which is automatically cleaned up
		// after the test completes.
		baseDir := t.TempDir()

		// Define the full path for the nested directory structure, including a file at the end.
		// nestedDir specifies a path with multiple levels of subdirectories.
		nestedDir := filepath.Join(baseDir, "level1", "level2", "level3", "file.txt")

		// Call RecursiveCreatePath to create the nested directory structure specified by nestedDir.
		// This function should create all intermediate directories and the final directory specified in the path.
		err := RecursiveCreatePath(nestedDir)
		// Assert that no error occurred during directory creation.
		// This ensures that the RecursiveCreatePath function executed successfully and all directories were created.
		assert.NoError(t, err, "Expected no error during directory creation")

		// Define the path to the deepest level of the created directory structure (excluding the file).
		// createdDir represents the path up to the last level of directories.
		createdDir := filepath.Join(baseDir, "level1", "level2", "level3")

		// Use os.Stat to check if the directory exists at the specified path.
		// This confirms that the intermediate directories were created as expected.
		_, err = os.Stat(createdDir)

		// Assert that the directory exists.
		// This ensures that the RecursiveCreatePath function created the entire directory hierarchy.
		assert.False(t, os.IsNotExist(err), "Expected directories to exist")

		// Remove the base directory and all its contents after the test.
		// This ensures no leftover files or directories remain from the test, maintaining a clean environment.
		_ = os.RemoveAll(baseDir)
	})

	// ExistingDirectory tests the behavior of the RecursiveCreatePath function when
	// the target directory already exists. The test verifies that the function does not
	// return an error if the directory is already present and that the directory remains unchanged.
	t.Run("ExistingDirectory", func(t *testing.T) {
		// Define the base directory where the nested directories will be created.
		// t.TempDir() creates a temporary directory for the test, which is automatically cleaned up
		// after the test completes.
		existingDir := t.TempDir()

		// Define the path to a file within the existing directory.
		// This path is used to test the RecursiveCreatePath function.
		existingFilePath := filepath.Join(existingDir, "file.txt")

		// Call RecursiveCreatePath with the path to a file within the existing directory.
		// The function should handle this gracefully, without attempting to create the file,
		// and without returning an error.
		err := RecursiveCreatePath(existingFilePath)

		// Verify that no error was returned by RecursiveCreatePath.
		// This confirms that the function handled the existing directory case correctly.
		assert.NoError(t, err, "Expected no error when directory already exists")

		// Use os.Stat to check the status of the directory specified by existingDir.
		// This function returns information about the file or directory, and an error if the file or directory does not exist.
		// Here, we are interested in confirming that the directory exists and no error occurred.
		_, err = os.Stat(existingDir)
		// Assert that no error occurred while checking the existence of the directory.
		// os.IsNotExist(err) returns true if the error indicates that the directory does not exist.
		// The assertion checks that this is not the case, meaning the directory should be present.
		// If the directory was not present, the test would fail, indicating that the RecursiveCreatePath function might have removed or failed to create it.
		assert.False(t, os.IsNotExist(err), "Expected existing directory to exist")

		// Remove the temporary directory and all its contents after the test.
		// This ensures no leftover files or directories remain from the test, maintaining a clean environment.
		_ = os.RemoveAll(existingDir)
	})
}
