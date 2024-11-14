package allocator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryAllocator(t *testing.T) {
	// Create an instance of the MemoryAllocator struct.
	// The MemoryAllocator is the type that implements the Allocator interface,
	// which provides methods for allocating and freeing memory.
	// By using the address-of operator (&), we ensure that alloc holds a pointer
	// to the instance of MemoryAllocator, allowing us to call methods on it later.
	alloc := &MemoryAllocator{}

	// MallocAllocatesMemory tests the Malloc method to ensure it correctly allocates memory.
	// This test checks that calling Malloc with a valid size returns a non-nil pointer.
	// The test also verifies that no error is returned, confirming that memory allocation was successful.
	// Finally, it ensures proper memory management by freeing the allocated memory to avoid leaks.
	t.Run("MallocAllocatesMemory", func(t *testing.T) {
		// Define the size of memory to be allocated, here 1024 bytes.
		// This value is chosen to test if Malloc can allocate a typical block of memory.
		size := 1024

		// Call the Malloc method to allocate memory of the specified size.
		// Malloc should return a non-nil pointer if allocation is successful.
		// If allocation fails, an error will be returned instead.
		ptr, err := alloc.Malloc(size)

		// Verify that no error occurred during the allocation.
		// The test will fail with the message "Malloc should succeed" if an error is returned.
		// This check ensures that the allocation process did not encounter any issues,
		// such as running out of memory or passing an invalid size.
		assert.NoError(t, err, "Malloc should succeed")

		// Verify that the pointer returned by Malloc is not nil, confirming successful allocation.
		// If the pointer is nil, it indicates that the memory allocation failed,
		// which would be an issue that needs to be addressed.
		assert.NotNil(t, ptr, "Malloc should return a non-nil pointer for successful allocation")

		// Free the allocated memory after the test to avoid memory leaks.
		// Calling Free releases the memory block back to the system, ensuring that
		// resources are properly cleaned up after the test completes.
		alloc.Free(ptr)
	})
}
