
package main

import (
    "os"
    "fmt"
    "flag"
    "inout"
)

/**
 * Usage
 */
func Usage() {
    fmt.Println("--------------------------------------------------------------------------------")
    fmt.Println("USAGE: system -action < ACTION >")
    fmt.Println("--------------------------------------------------------------------------------")

    fmt.Println("     encrypt -file FILE -key KEY [-name OUTPUT]")
    fmt.Println("     decrypt -file FILE -key KEY [-name OUTPUT]")
    fmt.Println("     rsa     [-size INTEGER]")
    fmt.Println("     ec      [-name <256 | 384 | 521>]")
    fmt.Println("     sign    -file FILE -key ECKEY [-name OUTPUT]")
    fmt.Println("     verify  -file FILE -key ECKEY  -sign SIGNATURE")
    fmt.Println("     send    -file FILE -key RSAKEY -sign ECKEY [-name OUTPUT]")
    fmt.Println("     receive -file FILE -key RSAKEY -sign ECKEY [-name OUTPUT] [-signName SIGNATURE]\n")

    fmt.Println("  * encrypt: Encrypts a file with a key usgin AES operating with CBC")
    fmt.Println("  * decrypt: Decrypts a file with a key usgin AES operating with CBC")
    fmt.Println("  * rsa:     Generates an RSA key in PEM format (public and private files). DEFAULT VALUE: 2048")
    fmt.Println("  * ec:      Generates an EC key in PEM format (public and private files). DEFAULT CURVE: P256")
    fmt.Println("  * sign:    Signs a file with the given key. Has to be EC")
    fmt.Println("  * verify:  Verifies if the file is signed with the signature given")
    fmt.Println("  * send:    Signs the message and encrypts it with AES including the key encrypted with RSA key")
    fmt.Println("  * receive: Decrypts the file and writes the message and the signature\n")

    fmt.Println("  -file FILE: Name of the file to use")
    fmt.Println("  -key  KEY:  Name of the key to use")
    fmt.Println("  -sign SIGN: Name of the key to sign/verify or signature")
    fmt.Println("  -name NAME: Name of the file to be written or curve to be used")
    fmt.Println("  -size SIZE: Number of bits to generate the RSA key")
    fmt.Println("  -signName NAME: Name of the file to be written as signature")
    fmt.Println("--------------------------------------------------------------------------------")
    os.Exit(0)
}

/** 
 * MAIN
 */
func main() {

    actionPtr := flag.String("action", "EMPTY", "Tell the action that you want")
    filePtr := flag.String("file", "FILE", "File to use")
    keyPtr  := flag.String("key", "KEY", "Key to use")
    signPtr := flag.String("sign", "SIGN", "Key to use for signing")
    namePtr := flag.String("name", "result.out", "Name of the file to write")
    signNamePtr := flag.String("signName", "result.signature", "Name of signature to write")
    sizePtr := flag.Int("size", 0, "Size")

    flag.Usage = func() {
        Usage()
    }
    flag.Parse()

    switch *actionPtr {
    default:
        Usage()

    case "encrypt":
        if *filePtr == "FILE" || *keyPtr == "KEY" {
            Usage()
        }
        var file []byte = inout.ReadFile(*filePtr)
        var key []byte = inout.ReadFile(*keyPtr) 
        fmt.Printf("\nEncrypting with key %s the file %s Output: (%s)\n",*filePtr,*keyPtr,*namePtr)

        var encrypted []byte = encrypt(file, key)
        inout.WriteFile(encrypted, *namePtr)

    case "decrypt": 
        if *filePtr == "FILE" || *keyPtr == "KEY" {
            Usage()
        }
        var file []byte = inout.ReadFile(*filePtr)
        var key []byte = inout.ReadFile(*keyPtr) 
        fmt.Printf("\nDecrypting with key %s the file %s Output: (%s)\n",*filePtr,*keyPtr,*namePtr)

        var decrypted []byte = decrypt(file, key)
        inout.WriteFile(decrypted, *namePtr)

    case "rsa":
        if *sizePtr == 0 {
            *sizePtr = 2048
        }
        fmt.Printf("\nGenerating RSA key of %d bytes.  Output: (publicRSA.pem, privateRSA.pem)\n",*sizePtr)
        RSAkey(*sizePtr)

    case "ec":
        if *namePtr == "result.out" {
            *namePtr = "256"
        }
        if *namePtr != "256" && *namePtr != "384" && *namePtr != "521" {
            Usage()
        } else {
            fmt.Printf("\nGenerating EC key with P%s curve. Output: (publicEC.pem, privateEC.pem)\n",*namePtr)
            ECCkey(*namePtr)
        }

    case "sign":
        if *filePtr == "FILE" || *keyPtr == "KEY" {
            Usage()
        }
        var file []byte = inout.ReadFile(*filePtr)
        var key []byte = inout.ReadFile(*keyPtr) 
        fmt.Printf("\nSigning with %s key the file %s Output: (%s)\n",*filePtr,*keyPtr,*namePtr)

        var signed []byte = signECDSA(file,key)
        inout.WriteFile(signed,*namePtr)
        
    case "verify":
        if *filePtr == "FILE" || *keyPtr == "KEY" || *signPtr == "SIGN" {
            Usage()
        }
        var file []byte = inout.ReadFile(*filePtr)
        var key []byte = inout.ReadFile(*keyPtr) 
        var sign []byte = inout.ReadFile(*signPtr)
        fmt.Printf("\nVerifying the file %s ...\n",*filePtr)

        if verifyECDSA(file,sign,key) {
            fmt.Printf("Verified correctly with the key %s\n",*keyPtr)
        } else {
            fmt.Printf("Sign doesn't match with the key %s\n",*keyPtr)
        }

    case "send":
        if *filePtr == "FILE" || *keyPtr == "KEY" || *signPtr == "SIGN" {
            Usage()
        }
        var file []byte = inout.ReadFile(*filePtr)
        var key []byte = inout.ReadFile(*keyPtr) 
        var sign []byte = inout.ReadFile(*signPtr)
        fmt.Printf("\nSigning with %s key the file %s, encrypting AES key with %s\n",*signPtr,*filePtr,*keyPtr)

        var result []byte = sendMessage(file,key,sign)
        fmt.Printf("\nOutput: %s\n", *namePtr)
        inout.WriteFile(result,*namePtr)

    case "receive":
        if *filePtr == "FILE" || *keyPtr == "KEY" || *signPtr == "SIGN" {
            Usage()
        }
        var file []byte = inout.ReadFile(*filePtr)
        var key []byte = inout.ReadFile(*keyPtr) 
        var sign []byte = inout.ReadFile(*signPtr)
        fmt.Printf("\nDecrypting with %s key AES key and decrypting and verifying \nwith %s the message in %s\n",*keyPtr,*signPtr,*filePtr)

        Mess, Firm := receiveMessage(file,key,sign)
        fmt.Printf("\nOutput: (%s, %s)\n", *namePtr, *signNamePtr)
        inout.WriteFile(Mess,*namePtr)
        inout.WriteFile(Firm,*signNamePtr)
    }
}
