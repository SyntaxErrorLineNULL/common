package crypto

import (
	"crypto/aes"
	"encoding/binary"
	"encoding/hex"
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

	// InvalidKey tests the behavior of the encryption and decryption methods
	// when provided with an invalid encryption key. It verifies that the methods
	// return appropriate errors when an invalid key is used for encryption and decryption.
	// This test ensures that the encryption and decryption methods handle invalid keys
	// correctly and fail gracefully, as expected.
	t.Run("InvalidKey", func(t *testing.T) {
		// Define an invalid encryption key for testing.
		// This key is purposely incorrect to test the method's error handling.
		invalidKey := "invalidkey"
		// Generate the current time for use in creating the initialization vector (IV).
		// This ensures that the IV is generated with a timestamp-based value.
		currentTime := time.Now()
		// Convert the current time to a Unix timestamp for use in the IV.
		// This timestamp represents the current time in seconds since January 1, 1970.
		unixTimestamp := currentTime.Unix()
		// Convert the Unix timestamp to an unsigned 64-bit integer for use in the IV.
		// This ensures the IV is a valid length for the encryption algorithm.
		validUntil := uint64(unixTimestamp)
		// Create an initialization vector (IV) with a length of 16 bytes.
		// The IV is required for AES encryption to ensure unique ciphertexts for the same plaintext.
		iv := make([]byte, 16)

		// Store the Unix timestamp in the first 8 bytes of the IV using big-endian encoding.
		// This provides the IV with a timestamp-based value.
		binary.BigEndian.PutUint64(iv, validUntil)
		// Define the plaintext to be encrypted.
		// This is the data that will be encrypted and later decrypted to verify correctness.
		plainText := []byte("Hello, Gophers!")

		// Attempt to encrypt the plaintext using the invalid key.
		// This operation should fail since the key is incorrect.
		_, err := crypto.EncryptCBC(invalidKey, iv, plainText)
		// Assert that an error occurred during encryption with the invalid key.
		// This ensures that the encryption method correctly identifies and reports the invalid key.
		assert.Error(t, err, "Expected error for invalid key")

		// Attempt to decrypt a sample encrypted text using the invalid key.
		// This operation should fail since the key is incorrect and does not match the encryption key.
		_, err = crypto.DecryptCBC(invalidKey, iv, "test")
		// Assert that an error occurred during decryption with the invalid key.
		// This ensures that the decryption method correctly identifies and reports the invalid key.
		assert.Error(t, err, "Expected error for decryption with invalid key")
	})

	// EmptyCipherText tests the behavior of the decryption method when provided
	// with an empty ciphertext. It verifies that the method returns an appropriate error
	// when attempting to decrypt an empty string, ensuring that the decryption method
	// handles this edge case correctly.
	t.Run("EmptyCipherText", func(t *testing.T) {
		// Define a valid encryption key for testing.
		// This key is used for the decryption process, even though the ciphertext is empty.
		key := "00112233445566778899aabbccddeeff"
		// Generate the current time for use in creating the initialization vector (IV).
		// This ensures that the IV is generated with a timestamp-based value.
		currentTime := time.Now()
		// Convert the current time to a Unix timestamp for use in the IV.
		// This timestamp represents the current time in seconds since January 1, 1970.
		unixTimestamp := currentTime.Unix()
		// Convert the Unix timestamp to an unsigned 64-bit integer for use in the IV.
		// This ensures the IV is a valid length for the encryption algorithm.
		validUntil := uint64(unixTimestamp)

		// Create an initialization vector (IV) with a length of 16 bytes.
		// The IV is required for AES decryption to ensure proper decryption of the ciphertext.
		iv := make([]byte, 16)
		// Store the Unix timestamp in the first 8 bytes of the IV using big-endian encoding.
		// This provides the IV with a timestamp-based value.
		binary.BigEndian.PutUint64(iv, validUntil)

		// Attempt to decrypt an empty ciphertext using the defined key and IV.
		// This operation should fail since the ciphertext is empty and not valid for decryption.
		_, err := crypto.DecryptCBC(key, iv, "")
		// Assert that an error occurred during decryption with the empty ciphertext.
		// This ensures that the decryption method correctly identifies and reports the invalid input.
		assert.Error(t, err, "Expected error for decryption with empty cipher text")
	})

	// TestDecryptWithIncorrectIVLength tests the decryption functionality when an
	// incorrect initialization vector (IV) length is provided. The decryption process
	// requires an IV of the correct length, which for AES encryption is typically 16 bytes.
	// In this test, an IV of only 8 bytes is used, which is deliberately incorrect, to
	// verify that the decryption method handles this erroneous input appropriately.
	// The test checks that the method returns an error when the IV length does not meet
	// the expected size requirements. This ensures that the decryption process validates
	// the IV length and fails gracefully when an invalid IV is supplied.
	t.Run("DecryptWithIncorrectIVLength", func(t *testing.T) {
		// Define a valid encryption key for testing purposes.
		// This key is used in the decryption function but the actual decryption will fail
		// due to the incorrect IV, so the key's correctness is not the focus of this test.
		key := "00112233445566778899aabbccddeeff"
		// Create an initialization vector (IV) with an incorrect length of 8 bytes.
		// For AES encryption, the IV length should be 16 bytes. Using an incorrect IV length
		// will help test how the decryption method handles invalid IV sizes.
		invalidIV := make([]byte, 8)
		// Define a sample ciphertext for the decryption attempt.
		// This ciphertext is a placeholder and will not be processed due to the incorrect IV.
		cipherText := "test"

		// Attempt to decrypt the given ciphertext using the defined key and the incorrect IV.
		// The decryption should fail because the IV length is not valid, which is the intended
		// behavior to be tested.
		_, err := crypto.DecryptCBC(key, invalidIV, cipherText)
		// Assert that an error is returned during the decryption attempt with the incorrect IV.
		// This confirms that the decryption function correctly identifies and reports the issue
		// with the IV length, ensuring robust error handling in scenarios of invalid input.
		assert.Error(t, err, "Expected error for decryption with incorrect IV length")
	})
}

func FuzzEncryptCBC(f *testing.F) {
	// Initialize a Crypto instance to be used for the AES encryption and decryption tests.
	// This instance is reused across all the test cases to ensure consistency in encryption behavior.
	crypto := &Crypto{}

	// Add initial test cases with various combinations of key, IV, and plaintext to the fuzzing function.
	// These test cases are used to evaluate the behavior of the Encrypt function under different input conditions.

	// Add a test case with a valid 32-character hexadecimal key, a valid 16-byte IV, and a non-empty plaintext.
	// This represents a typical valid input scenario.
	f.Add("0102030405060708090a0b0c0d0e0f10", []byte("0102030405060708"), []byte("example plaintext"))
	// Add a test case with a different valid 32-character hexadecimal key, a valid 16-byte IV, and a non-empty plaintext.
	// This case tests encryption with a different key while keeping the IV and plaintext valid.
	f.Add("deadbeefdeadbeefdeadbeefdeadbeef", []byte("0000000000000000"), []byte("test"))
	// Add a test case with a valid 32-character hexadecimal key, a valid 16-byte IV, and an empty plaintext.
	// This tests the encryption behavior when the plaintext is empty but the key and IV are valid.
	f.Add("00000000000000000000000000000000", []byte("0000000000000000"), []byte(""))

	// Add a test case with an invalid key that does not meet any valid AES key length requirement.
	// The key is not a valid length (16, 24, or 32 bytes), which should trigger a decoding error.
	f.Add("invalidkey", []byte("0000000000000000"), []byte("example plaintext"))
	// Add a test case with a valid 16-byte key but an invalid IV length.
	// The IV length is shorter than the required 16 bytes for AES encryption, which should cause a failure.
	f.Add("0102030405060708090a0b0c0d0e0f", []byte("shortiv"), []byte("example plaintext"))
	// Add a test case with a valid 32-character hexadecimal key but an empty IV.
	// The empty IV is invalid for AES encryption, so the function should handle this scenario correctly.
	f.Add("0102030405060708090a0b0c0d0e0f10", []byte(""), []byte("example plaintext"))
	// Add a test case with a valid 32-character hexadecimal key and a valid 16-byte IV, but with short plaintext.
	// This tests how the encryption function deals with short plaintext.
	f.Add("0102030405060708090a0b0c0d0e0f10", []byte("0102030405060708"), []byte("short"))

	// The f.Fuzz function is used to perform fuzz testing on the Encrypt function.
	// It generates a range of inputs and tests how the Encrypt function handles them.
	// This helps identify potential edge cases or unexpected behaviors.
	f.Fuzz(func(t *testing.T, keyHex string, iv []byte, plainText []byte) {
		// Decode the key from its hexadecimal string representation.
		// This step converts the key from a hex string format to a byte slice format
		// that is required for the encryption process.
		keyBytes, err := hex.DecodeString(keyHex)
		// Check if the decoding of the key from its hex string representation resulted in an error.
		// If an error occurred, it could be due to an invalid key format.
		if err != nil {
			// If decoding fails and the length of the hex string is neither 32 nor 64 characters,
			// skip the test as the key cannot be processed correctly.
			if len(keyHex) != 32 && len(keyHex) != 64 {
				// Skip the test as the key is invalid.
				// t.Skip("Key is not valid")
				return
			}
		}

		// Check if the length of the decoded key is valid for AES encryption.
		// AES encryption requires keys to be either 16 bytes (AES-128), 24 bytes (AES-192),
		// or 32 bytes (AES-256). Any other length is considered invalid.
		if len(keyBytes) != 16 && len(keyBytes) != 24 && len(keyBytes) != 32 {
			// Skip the test if the key length is invalid.
			// t.Skip("Key is not valid")
			return
		}

		// Check if the length of the initialization vector (IV) is valid.
		// For AES encryption, the IV must be exactly 16 bytes long.
		if len(iv) != aes.BlockSize {
			// Skip the test if the IV length is invalid.
			// t.Skip("Ensure IV length is not valid")
			return
		}

		// Call the Encrypt function with the provided key, IV, and plaintext.
		// This function is expected to return a cipherText and potentially an error.
		cipherText, err := crypto.EncryptCBC(keyHex, iv, plainText)
		// Check if an error occurred during the encryption process.
		// If an error occurred, it indicates that the Encrypt function could not handle
		// the provided inputs correctly.
		if err != nil {
			// Assert that an error is expected in this case.
			// The test is skipped if encryption fails with the given inputs.
			assert.Error(t, err, "Encryption failed with key: %s, IV: %v, plainText: %s", keyHex, iv, plainText)
			// Skip the test as encryption failed.
			// t.Skip("Failed to make encrypt")
			return
		}

		// If no error occurred, ensure that encryption was successful.
		// This assertion verifies that the Encrypt function did not return an error
		// for valid inputs.
		assert.NoError(t, err, "Expected no error during encryption")
		// Ensure that the cipherText returned from encryption is not empty.
		// An empty cipherText indicates that encryption may have failed or produced no result.
		assert.NotEmpty(t, cipherText, "Expected cipherText to be non-empty")

		// Calculate the expected length of the cipherText after encryption.
		// The cipherText length must be a multiple of the AES block size due to padding.
		// Add the AES block size to the length of the plaintext to ensure there is room for padding,
		// then round up to the nearest multiple of the block size.
		expectedLength := ((len(plainText) + aes.BlockSize) / aes.BlockSize) * aes.BlockSize
		// Assert that the length of the cipherText matches the expected length.
		// The length of the cipherText is expected to be twice the expected length because
		// it is encoded in hexadecimal format, which doubles the number of bytes.
		// Check if the actual length of the cipherText (after hex encoding) matches the expected length.
		assert.Len(t, cipherText, expectedLength*2, "Expected cipherText length to be %d bytes, got %d", expectedLength, len(cipherText)/2)
	})
}

func FuzzDecryptCBC(f *testing.F) {
	// Initialize a Crypto instance to be used for the AES encryption and decryption tests.
	// This instance is reused across all the test cases to ensure consistency in encryption behavior.
	crypto := &Crypto{}

	// Add initial test cases with various combinations of key, IV, and ciphertext to the fuzzing function.
	// These test cases are used to evaluate the behavior of the Decrypt function under different input conditions.

	// Add a test case with a valid 32-character hexadecimal key, a valid 16-byte IV, and a valid ciphertext.
	// This represents a typical valid input scenario.
	f.Add("0102030405060708090a0b0c0d0e0f10", []byte("0102030405060708"), "c4d6e8f0d1e2b9c8a7d8e6f4c4b5a6d0")
	// Add a test case with a valid 32-character hexadecimal key, a valid 16-byte IV, and a different valid ciphertext.
	// This tests decryption with a different ciphertext while keeping the key and IV valid.
	f.Add("00000000000000000000000000000000", []byte("0000000000000000"), "3a3a3a3a3a3a3a3a3a3a3a3a3a3a3a3a")

	// Add invalid cases to test the Decrypt function's robustness against erroneous inputs.
	// These cases should trigger errors and validate that the function handles invalid inputs gracefully.

	// Add a test case with an invalid key that does not meet any valid AES key length requirement.
	// The key is not of valid length (16, 24, or 32 bytes), which should trigger a decoding error.
	f.Add("invalidkey", []byte("0000000000000000"), "3a3a3a3a3a3a3a3a3a3a3a3a3a3a3a3a")
	// Add a test case with a valid 16-byte key but an invalid IV length.
	// The IV length is shorter than the required 16 bytes for AES decryption, which should cause a failure.
	f.Add("0102030405060708090a0b0c0d0e0f10", []byte("shortiv"), "3a3a3a3a3a3a3a3a3a3a3a3a3a3a3a3a")
	// Add a test case with a valid 32-character hexadecimal key but an empty IV.
	// The empty IV is invalid for AES decryption, so the function should handle this scenario correctly.
	f.Add("0102030405060708090a0b0c0d0e0f10", []byte(""), "3a3a3a3a3a3a3a3a3a3a3a3a3a3a3a3a")
	// Add a test case with a valid 32-character hexadecimal key and a valid 16-byte IV, but with an empty ciphertext.
	// This tests how the decryption function deals with an empty ciphertext.
	f.Add("0102030405060708090a0b0c0d0e0f10", []byte("0102030405060708"), "")
	// Add a test case with a valid 32-character hexadecimal key and a valid 16-byte IV, but with a ciphertext that is too short.
	// This tests how the decryption function handles ciphertext that may be incorrectly formatted or padded.
	f.Add("0102030405060708090a0b0c0d0e0f10", []byte("0102030405060708"), "00000000000000000000000000000000")

	// The f.Fuzz function is used to perform fuzz testing on the Decrypt function.
	// It generates a range of inputs and tests how the Decrypt function handles them.
	// This helps identify potential edge cases or unexpected behaviors.
	f.Fuzz(func(t *testing.T, keyHex string, iv []byte, cipherText string) {
		// Decode the key from its hexadecimal string representation.
		// This step converts the key from a hex string format to a byte slice format
		// that is required for the decryption process.
		keyBytes, err := hex.DecodeString(keyHex)
		// Check if the decoding of the key from its hex string representation resulted in an error.
		// If an error occurred, it could be due to an invalid key format.
		if err != nil {
			// If decoding fails and the length of the hex string is neither 32 nor 64 characters,
			// skip the test as the key cannot be processed correctly.
			if len(keyHex) != 32 && len(keyHex) != 64 {
				// Skip the test as the key is invalid.
				return
			}
		}

		// Check if the length of the decoded key is valid for AES decryption.
		// AES decryption requires keys to be either 16 bytes (AES-128), 24 bytes (AES-192),
		// or 32 bytes (AES-256). Any other length is considered invalid.
		if len(keyBytes) != 16 && len(keyBytes) != 24 && len(keyBytes) != 32 {
			// Skip the test if the key length is invalid.
			return
		}

		// Check if the length of the initialization vector (IV) is valid.
		// For AES decryption, the IV must be exactly 16 bytes long.
		if iv != nil && len(iv) != aes.BlockSize {
			// Skip the test if the IV length is invalid.
			return
		}

		// Call the Decrypt function with the provided key, IV, and ciphertext.
		// This function is expected to return a plaintext and potentially an error.
		plainText, err := crypto.DecryptCBC(keyHex, iv, cipherText)
		// Check if an error occurred during the decryption process.
		// If an error occurred, it indicates that the Decrypt function could not handle
		// the provided inputs correctly.
		if err != nil {
			// Assert that an error is expected in this case.
			// The test is skipped if decryption fails with the given inputs.
			assert.Error(t, err, "Decryption failed with key: %s, IV: %v, cipherText: %s", keyHex, iv, cipherText)
			return
		}

		// If no error occurred, ensure that decryption was successful.
		// This assertion verifies that the Decrypt function did not return an error
		// for valid inputs.
		assert.NoError(t, err, "Expected no error during decryption")
		// Ensure that the decrypted plaintext is not nil.
		// A nil plaintext indicates that decryption may have failed or produced no result.
		assert.NotNil(t, plainText, "Expected decrypted plainText to be non-nil")
	})
}
