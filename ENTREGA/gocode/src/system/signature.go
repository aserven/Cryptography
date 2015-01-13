package main

import (
    "errors"
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/ecdsa"
    "crypto/x509"
    "math/big"
    "encoding/pem"
)

/** 
 * Sign
 * Given the file to sign and the key in PEM format returns
 * the file signed with ECDSA.
 */
func signECDSA(file, key []byte) []byte {
    
    keyBlock, _ := pem.Decode(key)
    privatekey, err := x509.ParseECPrivateKey(keyBlock.Bytes)
    if err != nil {
        panic(err)
    }
    PSSmessage := file
    newhash := crypto.SHA256
    pssh := newhash.New()
    pssh.Write(PSSmessage)
    hashed := pssh.Sum(nil)

    r, s, error := ecdsa.Sign(rand.Reader, privatekey, hashed)
    if error != nil {
      panic(error)
    }

    return append(r.Bytes(),s.Bytes()...)
}

/** 
 * Verify
 * Given a file, that file signed and the verification key in 
 * PEM format, returns true if the signature with ECDSA is correct or false
 * if it is not.
 */
func verifyECDSA(file, signature, key []byte) bool {
    
    keyBlock, _ := pem.Decode(key)
    publickey, err := x509.ParsePKIXPublicKey(keyBlock.Bytes)
    if err != nil {
        panic(err)
    }
    PSSmessage := file
    newhash := crypto.SHA256
    pssh := newhash.New()
    pssh.Write(PSSmessage)
    hashed := pssh.Sum(nil)

    r := big.NewInt(0)
    s := big.NewInt(0)
    r = r.SetBytes(signature[:len(signature)/2])
    s = s.SetBytes(signature[len(signature)/2:])

    var succes bool
    if pkey, ok := publickey.(*ecdsa.PublicKey); ok {
        succes = ecdsa.Verify(pkey, hashed, r, s)
    } else {
        succes = false
    }
    if err != nil {
        panic(err)
    } else {
        return succes
    }
}



/******************************************************************************/
/***************************** RSA SIGN & VERIFICATION ************************/
/******************************************************************************/
 
/** 
 * Sign
 * Given the file to sign and the key in PEM format returns
 * the file signed with RSA.
 */
func signRSA(file, key []byte) []byte {

    keyBlock, _ := pem.Decode(key)

    privatekey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
    var opts rsa.PSSOptions
    opts.SaltLength = rsa.PSSSaltLengthAuto 

    PSSmessage := file
    newhash := crypto.SHA256
    pssh := newhash.New()
    pssh.Write(PSSmessage)
    hashed := pssh.Sum(nil)

    signaturePSS, err := rsa.SignPSS(rand.Reader, privatekey, newhash, hashed, &opts)
    if err != nil {
        panic(err)
    }

    return signaturePSS
}

/** 
 * Verify
 * Given a file, that file signed and the verification key in 
 * PEM format, returns true if the signature with RSA is correct or false
 * if it is not.
 */
 func verifyRSA(file, signature, key []byte) bool {

    keyBlock, _ := pem.Decode(key)
    publickey, err := x509.ParsePKIXPublicKey(keyBlock.Bytes)
    
    var opts rsa.PSSOptions
    opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

    PSSmessage := file
    newhash := crypto.SHA256
    pssh := newhash.New()
    pssh.Write(PSSmessage)
    hashed := pssh.Sum(nil)

    if pkey, ok := publickey.(*rsa.PublicKey); ok {
        err = rsa.VerifyPSS(pkey, newhash, hashed, signature, &opts)
    } else {
        err = errors.New("Public key is not an RSA key")
    }
    if err != nil {
     return false
    } else {
     return true
    }
 }