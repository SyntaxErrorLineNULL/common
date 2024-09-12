package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

// Crypto is an empty struct currently used as a placeholder or for future expansion.
// It may be utilized for cryptographic functions or settings related to encryption and decryption in the application.
type Crypto struct{}

// EncryptCBC performs AES encryption on the provided plaintext using the specified key and initialization vector (IV).
// It ensures the key, IV, and plaintext are valid before proceeding with the encryption. The key is decoded from a hexadecimal string,
// and padding is applied to the plaintext to meet the block size requirements for AES encryption. The method then encrypts the padded
// plaintext using AES in CBC mode with the given IV and returns the resulting ciphertext as a hexadecimal string. Any issues
// with the key, IV, or plaintext result in appropriate error messages.
func (srv *Crypto) EncryptCBC(key string, iv, plainText []byte) (string, error) {
	// Check for empty key, IV, or plaintext and return appropriate error messages.
	// These checks ensure that essential inputs are not missing.
	switch {
	case key == "", len(iv) == 0, len(plainText) == 0:
		return "", errors.New("key, IV block, or plaintext is empty")
	}

	// Decode the hexadecimal key string into a byte slice.
	// AES encryption requires the key to be in byte format, so we convert the provided hexadecimal string.
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher block using the decoded key.
	// The AES block will be used to encrypt the plaintext.
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		// Return an error if creating the cipher block fails
		return "", err
	}

	// Define the block size for AES encryption.
	blockSize := aes.BlockSize
	// Calculate the required padding for the plaintext.
	// AES uses block-based encryption, so the plaintext must be a multiple of the block size.
	// Padding ensures that the plaintext fits exactly into the blocks.
	padding := blockSize - len(plainText)%blockSize
	// Create a padding slice with the required padding value.
	// The padding value is the number of bytes needed to align plaintext with the block size.
	paddingBytes := bytes.Repeat([]byte{byte(padding)}, padding)
	// Append the padding bytes to the plaintext.
	// This ensures that the plaintext length is now a multiple of the AES block size, as required for proper encryption.
	plainText = append(plainText, paddingBytes...)

	// Create a slice to hold the encrypted data.
	// The length of this slice is the same as the padded plaintext length.
	cipherText := make([]byte, len(plainText))

	// Create a CBC mode encrypter using the AES block and the provided IV.
	// CBC (Cipher Block Chaining) mode requires an IV for encryption.
	mode := cipher.NewCBCEncrypter(block, iv)

	// Encrypt the padded plaintext using the CBC mode.
	// The result is stored in the cipherText byte slice.
	mode.CryptBlocks(cipherText, plainText)

	// Encode the resulting ciphertext into a hexadecimal string and return it.
	// The encrypted text is converted into a string format for easy storage or transmission.
	return hex.EncodeToString(cipherText), nil
}
