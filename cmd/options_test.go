package cmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	t.Parallel()

	// SetContext tests the behavior of the SetContext method in the Options struct.
	// The test ensures that the method correctly assigns a valid context to the
	// parentCtx field within the Options instance and returns no error.
	// This test verifies that the method handles a standard context (in this case,
	// context.Background()) appropriately and ensures that the context is stored
	// correctly for further use in operations that depend on it.
	t.Run("SetContext", func(t *testing.T) {
		// Initialize a new Options instance. This struct will be used to test
		// the SetContext method and confirm that the context is properly assigned.
		options := &Options{}

		// Define a valid context using context.Background().
		// This is a non-nil context that represents an empty, base context.
		ctx := context.Background()

		// Call the SetContext method, passing the valid context.
		// This should update the Options instance's parentCtx field with the provided context.
		err := options.SetContext(ctx)
		// Assert that no error was returned by the SetContext method.
		// This check confirms that the method successfully set the context without any issues.
		assert.NoError(t, err, "Expected no error when setting a valid context")
		// Assert that the context was correctly assigned to the parentCtx field.
		// This check ensures that the method stored the passed context properly
		// in the Options instance, making it available for future use.
		assert.Equal(t, ctx, options.parentCtx, "Expected parentCtx to be set to the provided context")
	})
}
