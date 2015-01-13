package main

import (
    "fmt"
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "errors"
    "encoding/pem"
)

/**
 * Send Message
 * Given a file, public RSA key and private EC key encrypts with AES the file with
 * its signing using the private EC key and the AES key encrypted with the public 
 * RSA key is added to the head
 */
func sendMessage(file, publicKeyRSA, signKeyEC []byte) []byte {
    
    var keyAES []byte = randomKey(16)
    var signed []byte = signECDSA(file,signKeyEC)
    var E_MF []byte = encrypt(append(file,signed...),keyAES)
    
    keyBlock, _ := pem.Decode(publicKeyRSA)
    publickey, err := x509.ParsePKIXPublicKey(keyBlock.Bytes)
    if err != nil {
        panic(err)
    }
    var KSE_EMF []byte
    if pkeyRSA, ok := publickey.(*rsa.PublicKey); ok {

        sha1hash := crypto.SHA1.New()
        encryptedKey, err := rsa.EncryptOAEP(sha1hash, rand.Reader, pkeyRSA, keyAES, nil)
        if err != nil {
           panic(err)
        }
        KSE_EMF = append(encryptedKey,E_MF...)

    } else {
        errors.New("Public key is not an RSA key")
    }
    return KSE_EMF
}

/**
 * Receive Message
 * Given an encrypted file, private RSA key and public EC key, decrypts the firsts
 * 2048 bytes with the private RSA key to reveal the AES key and then decrypts
 * the message with the signing and returns the pair
 */
func receiveMessage(file, privateKeyRSA, signKeyEC []byte) ([]byte, []byte) {
    
    var encKey []byte = file[:256]
    var encMess []byte = file[256:]

    keyBlock, _ := pem.Decode(privateKeyRSA)
    privatekey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
    if err != nil {
        panic(err)
    }
    sha1hash := crypto.SHA1.New()
    decryptedKey, err := rsa.DecryptOAEP(sha1hash, rand.Reader, privatekey, encKey, nil)
    if err != nil {
        panic(err)
    }
    var decryptedMsg []byte = decrypt(encMess,decryptedKey)

    var messageOK []byte = decryptedMsg[:len(decryptedMsg)-64]
    var signOK []byte = decryptedMsg[len(decryptedMsg)-64:]

    if verifyECDSA(messageOK,signOK,signKeyEC) {
        fmt.Println("Signed correctly: True")
    } else {
        fmt.Println("Signed correctly: False")
    }
    return messageOK, signOK
}


