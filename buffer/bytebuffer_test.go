package buffer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteBuffer(t *testing.T) {
	t.Parallel()

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

	// Reads data from an empty buffer and checks the behavior.
	// This test verifies that when attempting to read from an empty buffer,
	// the `Read` method returns an error and does not read any data. It ensures
	// that the buffer properly handles the empty state and responds correctly
	// with an error and zero bytes read.
	t.Run("ReadFromEmptyBuffer", func(t *testing.T) {
		// Create a new ByteBuffer instance without writing any data to it.
		// This ensures that the buffer is empty before the read operation, which
		// allows us to test how the Read method handles the empty buffer scenario.
		buf := &ByteBuffer{}

		// Prepare a buffer to read data into.
		// The read buffer is a byte slice of length 5, meaning it expects the Read
		// method to attempt reading 5 bytes. Since the buffer is empty, no data
		// will be written to this buffer.
		readBuffer := make([]byte, 5)

		// Attempt to read from the empty buffer.
		// Since the buffer is empty, we expect the Read method to return an error.
		// Additionally, no data should be read, so the number of bytes read should be 0.
		n, err := buf.Read(readBuffer)

		// Assert that an error occurred during the read operation.
		// The Read method should return an error because there is no data in the buffer.
		// The `assert.Error` function checks that the error is not nil, and if it is,
		// the test will fail with the message "Read should return an error when the buffer is empty".
		assert.Error(t, err, "Read should return an error when the buffer is empty")

		// Assert that no bytes were read from the empty buffer.
		// Since the buffer has no data, the number of bytes read should be 0.
		// The `assert.Equal` function checks that the number of bytes returned by the Read method is 0.
		// If the number of bytes is not 0, the test will fail with the message "No bytes should be read from an empty buffer".
		assert.Equal(t, 0, n, "No bytes should be read from an empty buffer")
	})

	// ReturnsCorrectBytes checks that data written to the buffer can be correctly retrieved.
	// This test ensures that the Write method appends data to the buffer without error,
	// and the Bytes method retrieves the exact data that was written. It confirms that
	// the buffer behaves correctly by storing and returning the expected bytes.
	t.Run("ReturnsCorrectBytes", func(t *testing.T) {
		// Create a new instance of ByteBuffer. This ensures that each test starts with a fresh buffer.
		// The buffer is initialized to be empty, so no previous data will interfere with the test.
		buf := &ByteBuffer{}

		// Define the data to be written to the buffer.
		// The byte slice contains the string "hello", which will be written to the buffer.
		// This data will be used to test the buffer's ability to store and return data correctly.
		data := []byte("hello")

		// Write data to the buffer using the Write method.
		// The Write method appends the provided data to the buffer and returns the number of bytes written.
		_, err := buf.Write(data)
		// Check that no error occurred during the Write operation.
		// The assert.NoError function ensures that the Write method executed without errors.
		// If an error occurs, the test will fail with the message "Write should not return an error".
		assert.NoError(t, err, "Write should not return an error")

		// Retrieve the bytes from the buffer using the Bytes method.
		// The Bytes method returns the byte slice that was written to the buffer.
		// This allows us to check if the buffer correctly stored the data written to it.
		bytes := buf.Bytes()
		// Verify that the bytes returned by the Bytes method are equal to the original data.
		// The assert.Equal function checks that the retrieved bytes match the data that was initially written.
		// If the data does not match, the test will fail with the message "Bytes should return the correct data".
		assert.Equal(t, data, bytes, "Bytes should return the correct data")
	})

	// ReadsDataCorrectly ensures that data can be read from a source into the buffer without errors.
	// This test checks that the ReadFrom method reads the correct number of bytes from the reader
	// into the buffer and that the buffer accurately stores the data read. It confirms the buffer's
	// ability to correctly store the data from the reader and reflects the exact content as expected.
	t.Run("ReadsDataCorrectly", func(t *testing.T) {
		// Initialize a new, empty ByteBuffer instance.
		// This buffer will hold the data read from the reader in the ReadFrom operation.
		buf := &ByteBuffer{}
		// Define a byte slice containing the test data "hello world".
		// This data will serve as the source content for reading into the ByteBuffer.
		data := []byte("hello world")
		// Create a new bytes.Reader to wrap the test data slice.
		// The bytes.Reader allows sequential reading of the data as required by the ReadFrom method.
		reader := bytes.NewReader(data)

		// Call ReadFrom to read data from the reader into the buffer.
		// This method reads the entire content of the reader and appends it to the buffer.
		// The return values include the number of bytes read and any error encountered.
		n, err := buf.ReadFrom(reader)
		// Assert that no error was returned during the read operation.
		// The assert.NoError function ensures the ReadFrom operation succeeded without errors.
		// If an error occurs, the test will fail with the message "ReadFrom should not return an error when reading valid data".
		assert.NoError(t, err, "ReadFrom should not return an error when reading valid data")
		// Assert that the number of bytes read matches the length of the data.
		// The assert.Equal function checks that the ReadFrom method returned the correct number of bytes.
		// If the byte count is incorrect, the test will fail with "ReadFrom should return the correct number of bytes read".
		assert.Equal(t, int64(len(data)), n, "ReadFrom should return the correct number of bytes read")
		// Assert that the buffer now contains the same data as the reader.
		// This check confirms that the buffer stored the data accurately by comparing it to the original data.
		// If the buffer content is incorrect, the test will fail with "Buffer should contain the same data as the reader after reading".
		assert.Equal(t, data, buf.Bytes(), "Buffer should contain the same data as the reader after reading")
	})

	// HandlesEmptyReader verifies the behavior of ByteBuffer when attempting to read from an empty reader.
	// This test ensures that the ReadFrom method does not produce an error when the input reader is empty,
	// returns zero bytes read, and leaves the ByteBuffer in an empty state. This behavior is important because
	// it confirms that the ReadFrom method can handle edge cases, such as reading from a source with no data,
	// without altering the buffer's contents or producing errors.
	t.Run("HandlesEmptyReader", func(t *testing.T) {
		// Create a new ByteBuffer instance that will serve as the target buffer for the read operation.
		// This buffer is initially empty, which allows us to observe any changes to it after attempting to read.
		buf := &ByteBuffer{}

		// Set up an empty reader using bytes.NewReader with an empty byte slice as input.
		// This reader simulates a data source with no content, which is critical to test how the ByteBuffer
		// handles the case of reading from a source that provides zero bytes of data.
		emptyReader := bytes.NewReader([]byte{})

		// Perform the read operation by calling ReadFrom with the empty reader.
		// This step attempts to read from the empty reader, and since there is no data, it should read zero bytes.
		// We will check both the byte count and error output to confirm that the method correctly handles empty input.
		n, err := buf.ReadFrom(emptyReader)

		// Confirm that no error occurred during the read operation.
		// Since the reader is empty, the ReadFrom method should not encounter any issues.
		// The assert.NoError function verifies that the error returned is nil, and if not, the test fails with
		// the message "ReadFrom should not return an error when reading from an empty reader."
		assert.NoError(t, err, "ReadFrom should not return an error when reading from an empty reader")

		// Check that the number of bytes read is zero, as expected for an empty reader.
		// Since there is no data in the reader, ReadFrom should report zero bytes read.
		// The assert.Equal function confirms that the number of bytes read matches the expected value of zero,
		// and if not, it fails with the message "ReadFrom should return 0 bytes read for an empty reader."
		assert.Equal(t, int64(0), n, "ReadFrom should return 0 bytes read for an empty reader")

		// Verify that the ByteBuffer remains empty after the read operation.
		// Given that the reader had no data, the buffer should not have any content added to it.
		// The assert.Equal function checks that the buffer's length is still zero, indicating no changes were made.
		// If the buffer is not empty, the test fails with the message "Buffer should remain empty after reading from an empty reader."
		assert.Equal(t, 0, buf.Len(), "Buffer should remain empty after reading from an empty reader")
	})
}
