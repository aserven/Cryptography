package main



import "crypto/rand"

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

 }


/** 
 * ECCkey
 * Given the name of the curve you want to generate 
 * returns two files in PEM format with the public and 
 * private key ECDH
 */
 func ECCkey() {
 	
 }

