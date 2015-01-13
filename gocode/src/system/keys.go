package main

import (
    "os"
    "crypto/rand"
    "crypto/rsa"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/x509"
    "encoding/pem"
)

/**
 * Generate random key
 * Given an integer n, generates a random key of length n
 */
func randomKey(n int) []byte {

    key := make([]byte, n)
    
    _, err := rand.Read(key)
    if err != nil {
        panic(err)
    }
    return key
}

/** 
 * RSA key
 * Given a number n returns two files in PEM format 
 * containing the public and private key RSA.
 */
func RSAkey(n int) {

    privatekey, err := rsa.GenerateKey(rand.Reader, n)
    if err != nil {
        panic(err)
    }

    var publickey *rsa.PublicKey = &privatekey.PublicKey
    publicPEMfile, err := os.Create("publicRSA.pem")
    if err != nil {
      panic(err)
    }
    publicEncodedKey, err := x509.MarshalPKIXPublicKey(publickey) 
    if err != nil {
      panic(err)
    }
    var publicPEMkey = &pem.Block{Type : "PUBLIC KEY", Bytes : publicEncodedKey}
    err = pem.Encode(publicPEMfile, publicPEMkey)
    if err != nil {
      panic(err)
    }
    publicPEMfile.Close()

    privatePEMfile, err := os.Create("privateRSA.pem")
    if err != nil {
      panic(err)
    }
    var privatePEMkey = &pem.Block{ Type : "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}
    err = pem.Encode(privatePEMfile, privatePEMkey)
    if err != nil {
      panic(err)
    }
    privatePEMfile.Close()
}

/** 
 * ECCkey
 * Given the name of the curve you want to generate 
 * returns two files in PEM format with the public and 
 * private key ECDH
 */
func ECCkey(name string) {

    var pubkeyCurve elliptic.Curve

    switch name {
    case "256": pubkeyCurve = elliptic.P256()
    case "384": pubkeyCurve = elliptic.P384()
    case "521": pubkeyCurve = elliptic.P521()
    }

    privatekey := new(ecdsa.PrivateKey)
    privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
    if err != nil {
        panic(err)
    }
    
    var publickey *ecdsa.PublicKey = &privatekey.PublicKey
    publicPEMfile, err := os.Create("publicEC.pem")
    if err != nil {
        panic(err)
    }
    publicEncodedKey, err := x509.MarshalPKIXPublicKey(publickey) 
    if err != nil {
        panic(err)
    }
    var publicPEMkey = &pem.Block{ Type : "PUBLIC KEY", Bytes: publicEncodedKey}
    err = pem.Encode(publicPEMfile, publicPEMkey)
    if err != nil {
        panic(err)
    }
    publicPEMfile.Close()

    privatePEMfile, err := os.Create("privateEC.pem")
    if err != nil {
        panic(err)
    }
    privateEncodedKey, err := x509.MarshalECPrivateKey(privatekey)
    var privatePEMkey = &pem.Block{ Type: "EC PRIVATE KEY", Bytes: privateEncodedKey}
    err = pem.Encode(privatePEMfile, privatePEMkey)
    if err != nil {
        panic(err)
    }
    privatePEMfile.Close()
}

