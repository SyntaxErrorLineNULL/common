# Crypto Package

This Go package provides basic AES encryption and decryption functionality using Cipher Block Chaining (CBC) mode. 
It includes two primary methods for working with AES encryption: EncryptCBC and DecryptCBC. 
These methods handle encryption and decryption using a provided key and initialization vector (IV).

## Overview

The package is centered around the Crypto struct, which currently serves as a placeholder for cryptographic functions. The two main methods allow encryption of plaintext into ciphertext and decryption of ciphertext back into plaintext using AES encryption in CBC mode.

## Features
* **AES Encryption (CBC Mode):** The EncryptCBC method provides encryption of plaintext using AES in CBC mode with a specified key and IV. It ensures the key and plaintext are valid and applies necessary padding to the plaintext before encryption.
* **AES Decryption (CBC Mode):** The DecryptCBC method decrypts ciphertext that was encrypted using AES in CBC mode. It validates the key, IV, and ciphertext and removes the padding applied during encryption to retrieve the original plaintext.

## Usage
#### Encrypting Plaintext

To encrypt plaintext, use the EncryptCBC method. This requires a hexadecimal-encoded key, a byte slice IV, and the plaintext to be encrypted.
```go
crypto := Crypto{}
cipherText, err := crypto.EncryptCBC("your-hexadecimal-key", iv, plainText)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Encrypted Text:", cipherText)
```

#### Decrypting Ciphertext

To decrypt ciphertext, use the DecryptCBC method. This requires the same hexadecimal-encoded key used for encryption, the same IV, and the ciphertext in hexadecimal format.

```go
crypto := Crypto{}
plainText, err := crypto.DecryptCBC("your-hexadecimal-key", iv, cipherText)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Decrypted Text:", string(plainText))
```

### Error Handling

Both methods validate the inputs and return descriptive errors if any issues are encountered, such as invalid key, IV, or ciphertext formats, or if the ciphertext size is incorrect. The encryption and decryption methods ensure secure processing by adhering to AES block size requirements.

## Requirements
* Go 1.17+

## Future Plans
The Crypto struct is currently a placeholder and may be expanded with additional cryptographic functions or configurations in future releases.
