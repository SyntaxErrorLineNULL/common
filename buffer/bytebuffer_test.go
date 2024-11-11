package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteBuffer(t *testing.T) {
	// Writes and reads data using the ByteBuffer struct.
	// This test ensures that data can be written to the buffer and read back correctly,
	// verifying the buffer's ability to store and retrieve data, as well as handling of empty buffer errors.
	t.Run("WriteAndReadData", func(t *testing.T) {
		// Create a new instance of ByteBuffer. This ensures that each test starts with a fresh buffer.
		buf := &ByteBuffer{}

		// Define the data to be written to the buffer.
		// This byte slice contains the string "hello", which will be written to the buffer in the next step.
		data := []byte("hello")

		// Write data to the buffer.
		// The Write method appends the data to the buffer and returns the number of bytes written and an error (if any).
		// We expect no error, so we use assert.NoError to check that the Write method behaves as expected.
		n, err := buf.Write(data)
		// Check that no error occurred during the Write operation.
		// If an error is returned, the test will fail with the message "Write should not return an error",
		// indicating that there was an unexpected issue when writing to the buffer.
		assert.NoError(t, err, "Write should not return an error")

		// Verify that the number of bytes written is equal to the length of the input data.
		// This ensures that the Write method is functioning correctly by appending the entire byte slice to the buffer.
		// The length of the data slice is expected to match the number of bytes written to the buffer.
		assert.Equal(t, len(data), n, "Write should return the correct number of bytes written")
		// Verify that the buffer's length is equal to the length of the data after writing.
		// After writing "hello", the length of the buffer should reflect the amount of data stored,
		// which in this case should be 5 bytes, the length of the string "hello".
		// This ensures that the Len method correctly reports the buffer's size after an operation.
		assert.Equal(t, len(data), buf.Len(), "Len should return the correct buffer size after writing")

		// Create a byte slice to hold the data that will be read from the buffer.
		// The slice is initialized to the same length as the original data that was written to the buffer.
		// This ensures that the buffer has enough space to store the data when it is read back.
		readBuffer := make([]byte, len(data))

		// Perform the Read operation to retrieve data from the buffer.
		// The Read method is expected to fill the readData slice with the bytes stored in the buffer.
		n, err = buf.Read(readBuffer)
		// Check that no error occurred during the Read operation.
		// If an error is returned, the test will fail with the message "Read should not return an error",
		// indicating that there was an unexpected issue when reading from the buffer.
		assert.NoError(t, err, "Read should not return an error")

		// Verify that the number of bytes read matches the length of the original written data.
		// The Read method should return the number of bytes that were read from the buffer.
		// We expect this number to match the length of the data that was originally written,
		// which in this case is the length of the byte slice `data`.
		assert.Equal(t, len(data), n, "Read should return the correct number of bytes read")

		// Verify that the data read from the buffer matches the original written data.
		// This is the key check to ensure that the buffer correctly stores and retrieves data.
		// The assertion checks that the byte slice readData contains exactly the same content
		// as the byte slice `data` that was originally written to the buffer.
		assert.Equal(t, data, readBuffer, "Read should return the correct data")
	})
}
