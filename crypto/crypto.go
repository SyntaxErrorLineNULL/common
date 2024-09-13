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

// DecryptCBC decrypts the given ciphertext using AES encryption with the specified key and initialization vector (IV).
// It checks the validity of the key, IV, and ciphertext, and then performs decryption using AES in CBC mode.
// The decrypted plaintext is returned after removing any padding added during encryption. If there are any issues
// with the key, IV, or ciphertext, appropriate error messages are returned.
func (srv *Crypto) DecryptCBC(key string, iv []byte, cipherText string) ([]byte, error) {
	// Check for empty key, IV, or ciphertext and return an appropriate error message.
	// These checks ensure that all required inputs are provided before attempting decryption.
	switch {
	case key == "", len(iv) == 0, cipherText == "":
		return nil, errors.New("key, IV block, or cipherText is empty")
	}

	// Decode the hexadecimal key string into a byte slice.
	// The AES decryption process requires the key to be in byte format.
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		// Return an error if decoding the key fails.
		return nil, err
	}

	// Decode the hexadecimal ciphertext string into bytes.
	// The ciphertext must be in byte format for decryption.
	cipherTextBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		// Return an error if decoding the ciphertext fails.
		return nil, err
	}

	// Create a new AES cipher block using the decoded key.
	// This block is used for decrypting the ciphertext.
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		// Return an error if creating the cipher block fails.
		return nil, err
	}

	// Check if the length of the ciphertext is a multiple of the AES block size.
	// AES encryption requires that the ciphertext length is a multiple of the block size.
	if len(cipherTextBytes)%aes.BlockSize != 0 {
		// Return an error if the ciphertext length is not a multiple of the block size.
		return nil, errors.New("cipherText is not a multiple of the block size")
	}

	// Create a CBC mode decrypter with the AES block and the provided IV.
	// CBC (Cipher Block Chaining) mode is used for decryption and requires an IV.
	mode := cipher.NewCBCDecrypter(block, iv)
	// Decrypt the ciphertext using CBC mode.
	// The decrypted data is written back into the cipherTextBytes slice.
	mode.CryptBlocks(cipherTextBytes, cipherTextBytes)

	// Ensure the decrypted ciphertext is not empty.
	// An empty result after decryption indicates an issue with the decryption process.
	if len(cipherTextBytes) == 0 {
		return nil, errors.New("cipherText is empty")
	}

	// Retrieve the padding value from the last byte of the decrypted data.
	// The padding value is used to determine how much padding was added during encryption.
	padding := int(cipherTextBytes[len(cipherTextBytes)-1])
	if padding < 1 || padding > aes.BlockSize {
		// Return an error if the padding value is invalid.
		return nil, errors.New("invalid padding size")
	}

	// Remove the padding from the decrypted data.
	// If padding is present, it is removed to retrieve the original plaintext.
	if padding != 0 {
		return cipherTextBytes[:len(cipherTextBytes)-padding], nil
	}

	// Return the decrypted plaintext as a byte slice.
	// If no padding is present, the plaintext is returned as is.
	return cipherTextBytes, nil
}
