

package main

import (
    "os"
    "fmt"
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "errors"
    "encoding/pem"
)



func sendMessage(file, keyRSA, signKey []byte) []byte {
    
    //var signed []byte = sign(file,key)
    var signed2 []byte = signECDSA(file,signKey)
    
    var keyAES []byte = randomKey(16)

    var M_F []byte = file
    var E_M_F []byte = encrypt(append(M_F,signed2...),keyAES)
    
    //var ret bool = verify(file,sign,key)
    /*var ret bool = verifyECDSA(file,signKey,keyECDS)
    if ret {
        fmt.Println("Signed correctly: True")
    } else {
        fmt.Println("Error signing: False")
    }*/

    // EncryptOAEP

    label := []byte("")
    sha1hash := crypto.SHA1.New()

    keyBlock2, _ := pem.Decode(keyRSA)

    fmt.Printf("KEY PEM: %s\n",keyBlock2.Type)

    publickey, _ := x509.ParsePKIXPublicKey(keyBlock2.Bytes)
    
    //var encryptedmsg []byte
    fmt.Printf("OAEP Encryption : \n")
    if pkeyRSA, ok := publickey.(*rsa.PublicKey); ok {
        /* act on pkey */
        //err = rsa.VerifyPSS(pkey, newhash, hashed, signature, &opts)
        encryptedmsg, err := rsa.EncryptOAEP(sha1hash, rand.Reader, pkeyRSA, keyAES, label)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }



        fmt.Printf("OAEP encrypted [%# x] to \n[%# x]\n", string(keyAES), encryptedmsg)
        fmt.Println(len(encryptedmsg))
        fmt.Printf("FINAL [%d]\n", len(append(encryptedmsg,E_M_F...)))
        return append(encryptedmsg,E_M_F...)
    } else {
        /* not a RSA public key */
        //encryptedmsg = make([]byte,2)
        errors.New("Public key is not an RSA key")
    }



    // DecryptOAEP
    /*
    decryptedmsg, err := rsa.DecryptOAEP(sha1hash, rand.Reader, privatekey, encryptedmsg, label)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Printf("OAEP decrypted [%x] to \n[%s]\n",encryptedmsg, decryptedmsg)
    fmt.Println()
    */
    return make([]byte,2)
}



func receiveMessage(file, keyRSA, signKey []byte) ([]byte, []byte) {
    
    //var signed []byte = sign(file,key)
    var encKey []byte = file[:256]
    
    var encMess []byte = file[256:]


    keyBlock, _ := pem.Decode(keyRSA)

    fmt.Printf("KEY PEM: %s\n",keyBlock.Type)

    privatekey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)

    //var ret bool = verify(file,sign,key)
    /*var ret bool = verifyECDSA(file,signKey,keyECDS)
    if ret {
        fmt.Println("Signed correctly: True")
    } else {
        fmt.Println("Error signing: False")
    }*/

    fmt.Printf("[%# x]\n",encKey)

    //label := []byte("")
    sha1hash := crypto.SHA1.New()

    // DecryptOAEP
    
    decryptedKey, err := rsa.DecryptOAEP(sha1hash, rand.Reader, privatekey, encKey, nil)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Printf("OAEP decrypted [%x] to \n[%# x]\n",encKey, decryptedKey)
    fmt.Println()

    var decryptedMsg []byte = decrypt(encMess,decryptedKey)

    var messageOK []byte = decryptedMsg[:len(decryptedMsg)-64]
    var signOK []byte = decryptedMsg[len(decryptedMsg)-64:]
    

    fmt.Printf("KEY PEM: %# x\n", messageOK)
     var ret bool = verifyECDSA(messageOK,signOK,signKey)
    if ret {
        fmt.Println("Signed correctly: True")
    } else {
        fmt.Println("Error signing: False")
    }
    return messageOK,signOK
}


