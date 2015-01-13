package main

import (
    "io"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "bytes"
    "encoding/binary"
)

/** 
 * pkcs7 Padding
 * Adds padding equal to blocksize specified
 */
func pkcs7PAD(file []byte, blocksize int) []byte {

    var p int = blocksize - len(file)%blocksize
    b := make([]byte, 2)
    binary.LittleEndian.PutUint16(b, uint16(p)) 
    b[1] = b[0]
    b = bytes.Repeat(b,p)

    return append(file,b[len(b)/2:]...)
}

/** 
 * pkcs7 UnPad
 * Removes padding specified by blocksize
 */
func pkcs7UNPAD(file []byte, blocksize int) []byte {

    b := make([]byte,2)
    b[0] = file[len(file)-1]
    b[1] = 0
    x := int(binary.LittleEndian.Uint16(b))

    return file[:len(file)-x] 
}

/** 
 * Encrypt
 * Given a file and a key, returns that file encrypted.
 * Using AES algorithm with 128 bits block size and CBC operator.
 */
func encrypt(file, key []byte) []byte {

    var plaintext []byte = pkcs7PAD(file,aes.BlockSize)

    if len(plaintext)%aes.BlockSize != 0 {
        panic("File is not a multiple of the block size")
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

    return ciphertext
}

/** 
 * Decrypt
 * Given an encrypted file and a key, returns that file decrypted.
 * Using AES algorithm with 128 bits with CBC operator.
 */
func decrypt(file, key []byte) []byte {

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    if len(file) < aes.BlockSize {
        panic("Ciphertext too short")
    }
    iv := file[:aes.BlockSize]
    ciphertext := file[aes.BlockSize:]

    if len(ciphertext)%aes.BlockSize != 0 {
        panic("Ciphertext is not a multiple of the block size")
    }
    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(ciphertext, ciphertext)
    decrypted := pkcs7UNPAD(ciphertext,aes.BlockSize)

    return decrypted
}
