package allocator

// #include <stdlib.h>
import "C"
import (
	"errors"
	"math"
	"unsafe"
)

// Allocator is an interface for managing memory allocation and deallocation.
// It abstracts the memory operations `Malloc` and `Free`, which are essential
// in scenarios requiring precise memory control, such as interacting with
// low-level C code or managing resources in a performance-critical environment.
// Implementers of Allocator can provide custom memory management, allowing
// for flexible handling of memory outside the Go garbage collector.
type Allocator interface {
	// Malloc allocates a block of memory of the specified size in bytes.
	// It returns a pointer to the allocated memory block if successful.
	// If the allocation fails, it returns an error, ensuring that the caller
	// is notified of any memory allocation issues immediately.
	Malloc(size int) (unsafe.Pointer, error)

	// Free releases the memory block at the provided pointer address.
	// By freeing this memory, it becomes available for future allocations,
	// which is critical in environments with limited resources or where
	// memory needs to be manually managed for efficiency.
	Free(pointer unsafe.Pointer)
}

// MemoryAllocator provides an implementation of the Allocator interface
// using C library functions to perform low-level memory operations. This
// structure enables allocation and deallocation of memory outside the Go
// garbage collector, which can be beneficial when working with large or
// time-sensitive allocations that require precise control.
type MemoryAllocator struct{}

// Malloc attempts to allocate a memory block of the specified size (in bytes).
// If the provided size is invalid (e.g., negative), it returns an error to
// prevent invalid memory allocations. The method uses C.malloc to perform
// the allocation and returns a pointer to the allocated block if successful.
func (alloc *MemoryAllocator) Malloc(size int) (unsafe.Pointer, error) {
	// Convert the size to a float64 to check if it's negative using math.Signbit.
	// The Signbit function is a robust way to verify if a value is negative,
	// even when size is cast to a different type for compatibility with C.
	if math.Signbit(float64(size)) {
		// Return an error message if the size is negative, as allocating
		// a negative block size is invalid and may cause undefined behavior.
		return nil, errors.New("size is negative")
	}

	// Call C.malloc with size cast to C.size_t, which is the expected type
	// for memory sizes in the C library. This ensures compatibility with
	// the C function and prevents potential issues from size mismatches.
	ptr := C.malloc(C.size_t(size))

	// Check if C.malloc returned a null pointer, which indicates that
	// memory allocation failed due to insufficient memory or other system limitations.
	// If ptr is nil, return an error message to notify the caller.
	if ptr == nil {
		return nil, errors.New("failed memory allocation")
	}

	// Return the pointer to the allocated memory block to the caller,
	// allowing them to use the memory as needed. The caller is responsible
	// for eventually freeing this memory to avoid leaks.
	return ptr, nil
}

// Free deallocates the memory block pointed to by the given pointer.
// By calling C.free, this method ensures that the allocated memory
// is released and made available for future allocations, helping prevent
// memory leaks and resource exhaustion. It is crucial that this method
// is called on all pointers returned by Malloc when they are no longer needed.
func (alloc *MemoryAllocator) Free(pointer unsafe.Pointer) {
	// Use C.free to release the memory at the provided pointer.
	// This step is essential to avoid memory leaks, as it informs the OS
	// that the memory can be reclaimed and reused. The Free method
	// only accepts valid pointers; freeing a null or invalid pointer
	// may cause undefined behavior.
	C.free(pointer)
}
