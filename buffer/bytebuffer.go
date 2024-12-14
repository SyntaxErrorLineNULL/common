package buffer

import (
	"errors"
	"io"
)

// ByteBuffer provides a simple buffer to store bytes in memory.
// This struct is intended to manage a dynamic byte slice that can be
// written to, read from, and queried for its current length.
// It includes methods for examining and modifying its contents, making
// it suitable for applications that need a temporary in-memory buffer.
type ByteBuffer struct {
	// Holds the underlying data in a byte slice, representing the contents of the buffer.
	// This slice dynamically grows as data is added, providing flexible storage capacity.
	bytes []byte
}

// Len returns the current length of the buffer in bytes.
// This method allows users to query the ByteBuffer for the amount of data it currently holds.
// The length is determined by the length of the internal byte slice, which grows as data is added.
func (b *ByteBuffer) Len() int {
	// Calculate the length of the byte slice by calling len on the bytes field.
	// The returned length represents the amount of data stored in the buffer.
	// This method is efficient and allows the caller to know the current buffer size.
	return len(b.bytes)
}

// Write appends the provided data to the buffer's byte slice.
// This method takes a byte slice as input, which represents the data to be added to the buffer.
// The function modifies the buffer's internal byte slice by appending the new data,
// extending the buffer's length accordingly, and ensuring the buffer stores all the written data.
// Write returns the new length of the buffer and a nil error, indicating the operation was successful.
func (b *ByteBuffer) Write(data []byte) (int, error) {
	// Append the input data to the buffer's internal byte slice.
	// The append function dynamically adjusts the buffer's capacity if needed,
	// allowing it to accommodate the added data without requiring explicit resizing.
	b.bytes = append(b.bytes, data...)

	// Return the updated length of the buffer as the number of bytes currently held.
	// This value is calculated by calling len on the updated byte slice,
	// providing the caller with the buffer's total size after the write operation.
	// Since no errors are expected during the append operation, nil is returned for the error.
	return len(b.bytes), nil
}

// Read reads data from the buffer into the provided byte slice p.
// This method transfers data from the buffer's internal byte slice to the provided slice,
// up to the length of p or the length of the buffer's data, whichever is smaller.
// If the buffer is empty, the method returns an error indicating no data is available.
// After reading, the buffer's internal slice is updated to exclude the data that has been read.
func (b *ByteBuffer) Read(p []byte) (int, error) {
	// Check if the buffer is empty by verifying the length of the bytes slice.
	// If the buffer has no data, return an error and zero bytes read.
	// This ensures the method signals an empty buffer state without attempting to copy data.
	if len(b.bytes) == 0 {
		return 0, errors.New("buffer is empty")
	}

	// Copy data from the buffer's internal byte slice to the provided slice p.
	// The copy function copies up to len(p) bytes or the length of the buffer, whichever is smaller.
	// This allows reading only a portion of the buffer if p has a smaller length, and ensures that
	// only available data is read, preventing out-of-bound access.
	n := copy(p, b.bytes)

	// Update the buffer's internal slice by removing the data that was read.
	// Slicing the byte slice from n onwards effectively discards the bytes that were copied,
	// maintaining only the unread data in the buffer.
	b.bytes = b.bytes[n:]

	// Return the number of bytes read and a nil error, as the operation was successful.
	return n, nil
}

// Bytes returns the current contents of the buffer as a byte slice.
// This method provides read-only access to the buffer's underlying data,
// allowing external functions to view or process the data stored in the buffer.
// The returned slice reflects the buffer’s state at the time of the call
// and does not modify the buffer itself.
func (b *ByteBuffer) Bytes() []byte {
	// Directly return the buffer's internal byte slice.
	// This gives external access to the data stored within the buffer without copying,
	// providing a lightweight way to read the buffer contents.
	return b.bytes
}

// ReadFrom reads data from the provided reader into the ByteBuffer. It continuously reads data from the reader,
// expanding the internal buffer when necessary, and keeps track of the total number of bytes read.
// The method returns the total number of bytes read and an error, if any. If the end of the stream is reached,
// it returns the total bytes read with a nil error. If another error occurs, it returns the total bytes read
// along with the encountered error. The buffer is dynamically resized to accommodate incoming data by doubling
// its capacity whenever it becomes full, ensuring efficient memory usage while reading potentially large amounts of data.
func (b *ByteBuffer) ReadFrom(reader io.Reader) (int64, error) {
	// Initialize a local variable `buffer` with the current byte slice stored in the ByteBuffer struct.
	// This variable allows manipulation of the buffer's contents as data is read from the reader.
	buffer := b.bytes
	// Capture the initial size of the buffer in bytes as an int64.
	// This value represents the current length of data already stored in the buffer
	// and will help calculate the total bytes read by the end of the function.
	initialSize := int64(len(buffer))
	// Determine the current capacity of the buffer, which represents the maximum number of bytes
	// that can be held in the buffer before a resize is necessary. This value is crucial for managing
	// buffer expansion efficiently as new data is read.
	bufferCapacity := int64(cap(buffer))
	// Initialize `totalBytesRead` with the value of `initialSize`.
	// This variable will keep a running count of all bytes read from the reader,
	// starting with the initial buffer length. It will be incremented as more data is read.
	totalBytesRead := initialSize

	// Check if the current buffer capacity is zero, which indicates that the buffer
	// is either nil or empty and therefore unable to hold any data.
	if bufferCapacity == 0 {
		// Set an initial capacity of 64 bytes for the buffer. This choice provides
		// a reasonable starting size to hold small amounts of data while allowing
		// for potential growth if more data is read.
		bufferCapacity = 64

		// Allocate a new byte slice with the specified initial capacity.
		// This buffer will store data as it is read from the reader.
		buffer = make([]byte, bufferCapacity)
	} else {
		// If the buffer already has some capacity, trim it to its current maximum capacity.
		// This step ensures that the buffer's length does not exceed its storage limit,
		// allowing data to be read and stored efficiently without exceeding its capacity.
		buffer = buffer[:bufferCapacity]
	}

	// Continuously reads from the reader until EOF or an error occurs.
	// If the buffer is full, it expands the buffer size to accommodate more data.
	// Tracks the total number of bytes read and handles errors like EOF or others.
	for {
		// Check if the total bytes read so far have reached the buffer's current capacity.
		// If so, it indicates that the buffer is full and needs to be expanded to accommodate more data.
		if totalBytesRead == bufferCapacity {
			// Double the current buffer capacity to create additional space for incoming data.
			// This growth strategy helps to minimize the number of reallocations,
			// providing an efficient way to handle increasing data volume.
			bufferCapacity *= 2

			// Create a new buffer with the updated (doubled) capacity.
			// This larger buffer will store both the data that has already been read
			// and any additional data that may be read from the reader.
			newBuffer := make([]byte, bufferCapacity)

			// Copy the contents of the current buffer into the newly created buffer.
			// This ensures that all previously read data is preserved within the expanded storage space.
			copy(newBuffer, buffer)

			// Assign the newly created, larger buffer to the buffer variable.
			// This step effectively replaces the original buffer with the expanded version,
			// preparing it to hold more data in subsequent read operations.
			buffer = newBuffer
		}

		// Attempt to read data from the reader into the buffer, starting at the position
		// immediately following the data already read (indexed by totalBytesRead).
		// The Read method will attempt to fill as much of the buffer as possible from this position onward.
		bytesRead, err := reader.Read(buffer[totalBytesRead:])

		// Update the total number of bytes read by adding the number of bytes just read in this operation.
		// This tracking is crucial for knowing how much valid data is currently stored in the buffer.
		totalBytesRead += int64(bytesRead)

		// Check if any error occurred during the read operation.
		// If an error exists, it may indicate the end of the stream (EOF) or another issue with reading.
		if err != nil {
			// Store the current buffer content up to the totalBytesRead position in the ByteBuffer struct.
			// This step ensures that any partially read data is preserved, even if an error occurred.
			b.bytes = buffer[:totalBytesRead]

			// Determine if the error is an EOF (End of File) error, indicating that the reader has
			// no more data to provide. In this case, return the total bytes read along with a nil error.
			if err == io.EOF {
				// EOF is not a fatal error; it just signals the end of data.
				return totalBytesRead, nil
			}

			// If a different error occurred, return the total bytes read so far along with the error.
			// This allows the caller to handle the error while still having access to any data that was read.
			return totalBytesRead, err
		}
	}
}

// Reset clears the content of the ByteBuffer by resetting its internal byte slice.
// It sets the length of the buffer to 0 while keeping the underlying capacity intact,
// effectively making the buffer empty without reallocating memory. This allows the buffer
// to be reused efficiently without the overhead of memory reallocation.
func (b *ByteBuffer) Reset() {
	// Resets the buffer's length to 0, effectively clearing its contents while keeping the underlying array intact.
	// This allows the buffer to be reused without reallocating memory.
	b.bytes = b.bytes[:0]
}
