package main


 /*import (
   "io"
   "os"
   "fmt"
   "crypto/rand"
   "crypto/rsa"
   "crypto/md5"
   "crypto"
 )
*/
import (
    "io"
    "os"
    "fmt"
    "hash"
    "math/big"

    "crypto/rand"
    "crypto/rsa"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/x509"

    "crypto/md5"

    "encoding/gob"
    "encoding/pem"
)

func randomKey(n int) []byte {
    
    key := make([]byte, n)

    _, err := rand.Read(key)
    if err != nil {
        // handle error here
    }
    return key
}

/** 
 * RSA key
 * Given a number n returns two files in PEM format 
 * containing the public and private key RSA.
 */
 func RSAkey(n int) {

   // generate private key
   privatekey, err := rsa.GenerateKey(rand.Reader, n)

   if err != nil {
     fmt.Println(err.Error)
     os.Exit(1)
   }

   D := privatekey.D //private exponent
   Primes := privatekey.Primes
   PCValues := privatekey.Precomputed

   // Note : Only used for 3rd and subsequent primes
   //CRTVal := privatekey.Precomputed.CRTValues


   fmt.Println("Private Key : ", privatekey)
   fmt.Println("\nPrivate Exponent : ", D.String())
   fmt.Printf("\nPrimes :\n %s\n %s \n", Primes[0].String(), Primes[1].String())
   fmt.Printf("\nPrecomputed Values :\n Dp[%s]\n Dq[%s]\n", PCValues.Dp.String(), PCValues.Dq.String())
   fmt.Printf("\nPrecomputed Values : Qinv[%s]", PCValues.Qinv.String())
   fmt.Println()

   var publickey *rsa.PublicKey
   publickey = &privatekey.PublicKey

   // save private and public key separately
   privatekeyfile, err := os.Create("private.key")
   if err != nil {
    fmt.Println(err)
    os.Exit(1)
   }
   privatekeyencoder := gob.NewEncoder(privatekeyfile)
   privatekeyencoder.Encode(privatekey)
   privatekeyfile.Close()


   publickeyfile, err := os.Create("public.key")
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   publickeyencoder := gob.NewEncoder(publickeyfile)
   publickeyencoder.Encode(publickey)
   publickeyfile.Close()

   // save PEM file
   publicPEMfile, err := os.Create("public.pem")

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   publicEncodedKey, err := x509.MarshalPKIXPublicKey(publickey) 
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   // http://golang.org/pkg/encoding/pem/#Block
   var publicPEMkey = &pem.Block{
                Type : "PUBLIC KEY",
                Bytes : publicEncodedKey}

   err = pem.Encode(publicPEMfile, publicPEMkey)

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   publicPEMfile.Close()
   

   privatePEMfile, err := os.Create("private.pem")

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   // http://golang.org/pkg/encoding/pem/#Block
   var privatePEMkey = &pem.Block{
                Type : "RSA PRIVATE KEY",
                Bytes : x509.MarshalPKCS1PrivateKey(privatekey)}

   err = pem.Encode(privatePEMfile, privatePEMkey)

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   privatePEMfile.Close()

 }

  /*

   // generate private key
   privatekey, err := rsa.GenerateKey(rand.Reader, n)

   if err != nil {
     fmt.Println(err.Error)
     os.Exit(1)
   }

   D := privatekey.D //private exponent
   Primes := privatekey.Primes
   PCValues := privatekey.Precomputed

   // Note : Only used for 3rd and subsequent primes
   //CRTVal := privatekey.Precomputed.CRTValues


   fmt.Println("Private Key : ", privatekey)
   fmt.Println()
   fmt.Println("Private Exponent : ", D.String())
   fmt.Println()
   fmt.Printf("Primes :\n %s\n %s \n", Primes[0].String(), Primes[1].String())
   fmt.Println()
   fmt.Printf("Precomputed Values :\n Dp[%s]\n Dq[%s]\n", PCValues.Dp.String(), PCValues.Dq.String())
   fmt.Println()
   fmt.Printf("Precomputed Values : Qinv[%s]", PCValues.Qinv.String())
   fmt.Println()



   // Note : Only used for 3rd and subsequent primes
   //fmt.Printf("CRTValues : Exp[%s]\n Coeff[%s]\n R[%s]\n", CRTVal[2].Exp.String(), CRTVal[2].Coeff.String(), CRTVal[2].R.String())

   // Note : if you want to have multi primes,
   //        use rsa.GenerateMultiPrimeKey() function instead of
   //        rsa.GenerateKey() function
   //        see http://golang.org/pkg/crypto/rsa/#GenerateMultiPrimeKey

   // http://golang.org/pkg/crypto/rsa/#PrivateKey.Precompute
   privatekey.Precompute()

   // http://golang.org/pkg/crypto/rsa/#PrivateKey.Validate
   err = privatekey.Validate()

   if err != nil {
     fmt.Println(err.Error)
     os.Exit(1)
   }

   var publickey *rsa.PublicKey
   publickey = &privatekey.PublicKey
   N := publickey.N // modulus
   E := publickey.E // public exponent

   fmt.Println("Public key ", publickey)
   fmt.Println()
   fmt.Println("Public Exponent : ", E)
   fmt.Println()
   fmt.Println("Modulus : ", N.String())
   fmt.Println()

   // EncryptOAEP
   msg := []byte("The secret message!")
   label := []byte("")
   md5hash := md5.New()

   encryptedmsg, err := rsa.EncryptOAEP(md5hash, rand.Reader, publickey, msg, label)

   if err != nil {
     fmt.Println(err)
     os.Exit(1)
   }

   fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(msg), encryptedmsg)
   fmt.Println()


   // DecryptOAEP
   decryptedmsg, err := rsa.DecryptOAEP(md5hash, rand.Reader, privatekey, encryptedmsg, label)

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   fmt.Printf("OAEP decrypted [%x] to \n[%s]\n",encryptedmsg, decryptedmsg)
   fmt.Println()


   // EncryptPKCS1v15
   encryptedPKCS1v15, errPKCS1v15 := rsa.EncryptPKCS1v15(rand.Reader, publickey, msg)

  if errPKCS1v15 != nil {
     fmt.Println(errPKCS1v15)
     os.Exit(1)
   }

   fmt.Printf("PKCS1v15 encrypted [%s] to \n[%x]\n", string(msg), encryptedPKCS1v15)
   fmt.Println()

   // DecryptPKCS1v15
   decryptedPKCS1v15, err := rsa.DecryptPKCS1v15(rand.Reader, privatekey, encryptedPKCS1v15)

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   fmt.Printf("PKCS1v15 decrypted [%x] to \n[%s]\n",encryptedPKCS1v15, decryptedPKCS1v15)
   fmt.Println()


   // SignPKCS1v15
   var h crypto.Hash
   message := []byte("This is the message to be signed!")
   hash := md5.New()
   io.WriteString(hash, string(message))
   hashed := hash.Sum(nil)

   signature, err := rsa.SignPKCS1v15(rand.Reader, privatekey,h, hashed)

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   fmt.Printf("PKCS1v15 Signature : %x\n", signature)

   //VerifyPKCS1v15
   err = rsa.VerifyPKCS1v15(publickey, h, hashed, signature)

   if err != nil {
     fmt.Println("VerifyPKCS1v15 failed")
     os.Exit(1)
   } else {
     fmt.Println("VerifyPKCS1v15 successful")
   }
   fmt.Println()

   // SignPSS
   var opts rsa.PSSOptions
   opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

   PSSmessage := []byte("Message to be PSSed!")
  newhash := crypto.MD5
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

 }*/


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

    //pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256


    privatekey := new(ecdsa.PrivateKey)
    privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    var publickey *ecdsa.PublicKey
    publickey = &privatekey.PublicKey

    fmt.Println("Private Key :")
    fmt.Printf("%x \n", privatekey)

    fmt.Println("Public Key :")
    fmt.Printf("%x \n",publickey)


   // save private and public key separately
   privatekeyfile, err := os.Create("privateEC.key")
   if err != nil {
    fmt.Println(err)
    os.Exit(1)
   }
   privatekeyencoder := gob.NewEncoder(privatekeyfile)
   privatekeyencoder.Encode(privatekey)
   privatekeyfile.Close()


   publickeyfile, err := os.Create("publicEC.key")
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   publickeyencoder := gob.NewEncoder(publickeyfile)
   publickeyencoder.Encode(publickey)
   publickeyfile.Close()

   // save PEM file
   publicPEMfile, err := os.Create("publicEC.pem")

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   publicEncodedKey, err := x509.MarshalPKIXPublicKey(publickey) 
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   // http://golang.org/pkg/encoding/pem/#Block
   var publicPEMkey = &pem.Block{
                Type : "PUBLIC KEY",
                Bytes : publicEncodedKey}

   err = pem.Encode(publicPEMfile, publicPEMkey)

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   publicPEMfile.Close()
   

   privatePEMfile, err := os.Create("privateEC.pem")

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   privateEncodedKey, err := x509.MarshalECPrivateKey(privatekey)
   // http://golang.org/pkg/encoding/pem/#Block
   var privatePEMkey = &pem.Block{
                Type : "EC PRIVATE KEY",
                Bytes : privateEncodedKey}

   err = pem.Encode(privatePEMfile, privatePEMkey)

   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   privatePEMfile.Close()

 

    // Sign ecdsa style

    var h hash.Hash
    h = md5.New()
    r := big.NewInt(0)
    s := big.NewInt(0)

    io.WriteString(h, "This is a message to be signed and verified by ECDSA!")
    signhash := h.Sum(nil)

    r, s, serr := ecdsa.Sign(rand.Reader, privatekey, signhash)
    if serr != nil {
       fmt.Println(err)
       os.Exit(1)
    }

    signature := r.Bytes()
    signature = append(signature, s.Bytes()...)

    fmt.Printf("Signature : %x\n", signature)

    // Verify
    verifystatus := ecdsa.Verify(publickey, signhash, r, s)
    fmt.Println(verifystatus)  // should be true
 }

