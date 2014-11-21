package main


import (
    "fmt"
    "hash"
    "crypto/rsa"
    "crypto/sha256"
)

/** 
 * Sign
 * Given the file to sign and the key in PEM format returns
 * the file signed.
 */
 func sign(file, key []byte) []byte {


 // SignPSS
   var opts rsa.PSSOptions
   opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

   PSSmessage := file
   newhash := crypto.sha256
   pssh := newhash.New()
   pssh.Write(PSSmessage)
   hashed = pssh.Sum(nil)


   signaturePSS, err := rsa.SignPSS(rand.Reader, privatekey, newhash, hashed, &opts)

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   fmt.Printf("PSS Signature : %x\n", signaturePSS)

   //VerifyPSS
   err = rsa.VerifyPSS(publickey, newhash, hashed, signaturePSS, &opts)

   if err != nil {
     fmt.Println("VerifyPSS failed")
     os.Exit(1)
   } else {
     fmt.Println("VerifyPSS successful")
   }

 }


/** 
 * Verify
 * Given a file, that file signed and the verification key in 
 * PEM format, returns true if the signature is correct or false
 * if it is not.
 */
 func verify() {
 	
 }