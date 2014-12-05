package main


import (
    "os"
    "fmt"
    //"hash"
    "errors"
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/ecdsa"
    //"crypto/sha256"
    "crypto/x509"

    "math/big"
    "encoding/pem"

)

/** 
 * Sign
 * Given the file to sign and the key in PEM format returns
 * the file signed.
 */
func sign(file, key []byte) []byte {

    keyBlock, _ := pem.Decode(key)

    fmt.Printf("KEY PEM: %s\n",keyBlock.Type)

    privatekey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
    // SignPSS
    var opts rsa.PSSOptions
    opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

    PSSmessage := file
    newhash := crypto.SHA256
    pssh := newhash.New()
    pssh.Write(PSSmessage)
    hashed := pssh.Sum(nil)


    signaturePSS, err := rsa.SignPSS(rand.Reader, privatekey, newhash, hashed, &opts)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Printf("PSS Signature : %x\n", signaturePSS)

    //VerifyPSS
    /*
    err = rsa.VerifyPSS(publickey, newhash, hashed, signaturePSS, &opts)

    if err != nil {
     fmt.Println("VerifyPSS failed")
     os.Exit(1)
    } else {
     fmt.Println("VerifyPSS successful")
    }*/
    return signaturePSS

}

func signECDSA(file, key []byte) []byte {
    keyBlock, _ := pem.Decode(key)

    fmt.Printf("KEY PEM: %s\n",keyBlock.Type)

    privatekey, err := x509.ParseECPrivateKey(keyBlock.Bytes)
    // SignPSS
    //var opts rsa.PSSOptions
    //opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

    PSSmessage := file
    newhash := crypto.SHA256
    pssh := newhash.New()
    pssh.Write(PSSmessage)
    hashed := pssh.Sum(nil)


    r, s, error := ecdsa.Sign(rand.Reader, privatekey, hashed)

    if error != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    var signatureEC []byte = append(r.Bytes(),s.Bytes()...)

    fmt.Println(r)
    fmt.Println(s)

    fmt.Printf("PSS Signature : %x\n%x\n", r, s)

    //VerifyPSS
    /*
    err = rsa.VerifyPSS(publickey, newhash, hashed, signaturePSS, &opts)

    if err != nil {
     fmt.Println("VerifyPSS failed")
     os.Exit(1)
    } else {
     fmt.Println("VerifyPSS successful")
    }*/
    return signatureEC
}


/** 
 * Verify
 * Given a file, that file signed and the verification key in 
 * PEM format, returns true if the signature is correct or false
 * if it is not.
 */
 func verify(file, signature, key []byte) bool {

    keyBlock2, _ := pem.Decode(key)

    fmt.Printf("KEY PEM: %s\n",keyBlock2.Type)

    publickey, err := x509.ParsePKIXPublicKey(keyBlock2.Bytes)
    
    //if publickey, ok := x509.ParsePKIXPublicKey(keyBlock.Bytes); ok {
        /* act on str */
    

    // SignPSS
    var opts rsa.PSSOptions
    opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

    PSSmessage := file
    newhash := crypto.SHA256
    pssh := newhash.New()
    pssh.Write(PSSmessage)
    hashed := pssh.Sum(nil)

    /*signaturePSS, err := rsa.SignPSS(rand.Reader, publickey, newhash, hashed, &opts)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    
    fmt.Printf("PSS Signature : %x\n", signaturePSS)*/
    if pkey, ok := publickey.(*rsa.PublicKey); ok {
        /* act on pkey */
        err = rsa.VerifyPSS(pkey, newhash, hashed, signature, &opts)
    } else {
        /* not a RSA public key */
        err = errors.New("Public key is not an RSA key")
    }
    //VerifyPSS
    //err = rsa.VerifyPSS(publickey, newhash, hashed, file, &opts)

    if err != nil {
     fmt.Println("VerifyPSS failed")
     return false
    } else {
     fmt.Println("VerifyPSS successful")
     return true
    }
    /*} else {
        /* not string *
        fmt.Println("VerifyPSS failed")
        os.Exit(1)
    }*/	
 }

func verifyECDSA(file, signature, key []byte) bool {
    keyBlock2, _ := pem.Decode(key)

    fmt.Printf("KEY PEM: %s\n",keyBlock2.Type)

    publickey, err := x509.ParsePKIXPublicKey(keyBlock2.Bytes)
    
    //if publickey, ok := x509.ParsePKIXPublicKey(keyBlock.Bytes); ok {
        /* act on str */
    

    // SignPSS
    var opts rsa.PSSOptions
    opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

    PSSmessage := file
    newhash := crypto.SHA256
    pssh := newhash.New()
    pssh.Write(PSSmessage)
    hashed := pssh.Sum(nil)

    r := big.NewInt(0)
    s := big.NewInt(0)

    r = r.SetBytes(signature[:len(signature)/2])
    s = s.SetBytes(signature[len(signature)/2:])

    fmt.Printf("length %d\n", len(signature))
    fmt.Printf("numer R length %d\n %s\n ", len(signature)/2,r)
    fmt.Printf("numer S length %d\n %s\n ", len(signature)/2,s)
    /*signaturePSS, err := rsa.SignPSS(rand.Reader, publickey, newhash, hashed, &opts)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    /AGAFAR BYTES I PASAR A BIG INT
    
    fmt.Printf("PSS Signature : %x\n", signaturePSS)*/
    var succes bool
    if pkey, ok := publickey.(*ecdsa.PublicKey); ok {
        /* act on pkey */
        succes = ecdsa.Verify(pkey, hashed, r, s)
    } else {
        /* not a RSA public key */
        //err = errors.New("Public key is not an ECDSA key")
        succes = false
    }
    //VerifyPSS
    //err = rsa.VerifyPSS(publickey, newhash, hashed, file, &opts)

    if err != nil {
     fmt.Println("VerifyPSS ec failed")
     return false
    } else {
     fmt.Println("VerifyPSS ec successful")
     return succes
    }
    /*} else {
        /* not string *
        fmt.Println("VerifyPSS failed")
        os.Exit(1)
    }*/ 
}