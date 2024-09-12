package crypto

import (
	"encoding/binary"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {
	t.Parallel()

	// Initialize a Crypto instance to be used for the AES encryption and decryption tests.
	// This instance is reused across all the test cases to ensure consistency in encryption behavior.
	crypto := &Crypto{}

	// Encrypts and decrypts data using AES with a valid key and initialization vector (IV).
	// This test verifies that the encryption and decryption processes work correctly with
	// the given key and IV. It ensures that the encrypted data can be decrypted back to
	// the original plain text, confirming the correctness of the AES implementation.
	t.Run("EncryptsAndDecrypts", func(t *testing.T) {
		// Capture the current time to be used for generating the initialization vector (IV).
		// This ensures that the timestamp used is accurate and reflects the current moment.
		currentTime := time.Now()
		// Convert the current time to a Unix timestamp (seconds since January 1, 1970).
		// This timestamp is used to provide a unique value for the initialization vector (IV).
		unixTimestamp := currentTime.Unix()
		// Cast the Unix timestamp to a uint64 value for use in the IV.
		// The uint64 type is used because the IV for AES-128 requires 16 bytes,
		// and a 64-bit integer provides sufficient entropy for the IV.
		validUntil := uint64(unixTimestamp)
		// Create a byte slice of length 16 to serve as the initialization vector (IV) for AES-128.
		// The length of the IV is specific to the AES encryption algorithm, where 16 bytes are required.
		iv := make([]byte, 16)
		// Encode the uint64 timestamp into the first 8 bytes of the IV using big-endian byte order.
		// This ensures the timestamp is correctly represented within the IV and is compatible with AES encryption.
		binary.BigEndian.PutUint64(iv, validUntil)

		// Define the AES encryption key. This is a hexadecimal string representing a 128-bit key.
		// AES encryption requires a secret key, and in this case, a fixed key is used for testing.
		key := "00112233445566778899aabbccddeeff"
		// Define the plain text that will be encrypted. This text is a simple byte slice containing
		// a message. After encryption, this will be transformed into ciphertext, which should be
		// reversible through decryption back to the original plain text.
		plainText := []byte("Hello, Gophers!")

		// Encrypt the plain text using the specified AES key and initialization vector (IV).
		// The `crypto.Encrypt` function performs the encryption operation and returns the encrypted data.
		encrypted, err := crypto.EncryptCBC(key, iv, plainText)
		// Check if an error occurred during the encryption process.
		// The `assert.NoError` function verifies that the error is nil, meaning the encryption was successful.
		// If an error is present, the test will fail with the message "Encryption failed",
		// indicating that the encryption operation did not complete as expected.
		assert.NoError(t, err, "Encryption failed")

		// Decrypt the previously encrypted data using the same AES key and initialization vector (IV) that were used for encryption.
		// The `crypto.Decrypt` function performs the decryption operation and returns the decrypted data.
		decrypted, err := crypto.DecryptCBC(key, iv, encrypted)
		// Check if an error occurred during the decryption process.
		// The `assert.NoError` function verifies that the error is nil, meaning the decryption was successful.
		// If an error is present, the test will fail with the message "Decryption failed",
		// indicating that the decryption operation did not complete as expected.
		assert.NoError(t, err, "Decryption failed")
		// Verify that the decrypted data matches the original plain text.
		// The `assert.Equal` function checks if the decrypted data is equal to the plain text that was initially encrypted.
		// If the decrypted data does not match the original plain text, the test will fail with the message "Decrypted text does not match original",
		// indicating that the decryption process did not restore the original data correctly.
		assert.Equal(t, plainText, decrypted, "Decrypted text does not match original")
	})
}
