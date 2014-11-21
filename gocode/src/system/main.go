
package main

import (
    "os"
    "fmt"
    "flag"
    "strings"
    "inout"
)

func Usage() {
    // PER FER ...
}


// MIRAR INTERFACES I CHANNELS
// CAMBIAR A ACTION = [encrypt, decrypt, rsa, genkey] 

func main() {

    filePtr := flag.String("file", "FILE", "File to encrypt/decrypt")
    keyPtr  := flag.String("key", "KEY", "Key to encrypt/decrypt")
    namePtr := flag.String("name", "NAME", "Name of the file to write")
    rsaPtr := flag.Int("rsa", 2048, "Generate RSA key of n bits")
    eccPtr := flag.String("ECC", "256", "Generate ECC key called P256, P384 or P521")
    encPtr := flag.Bool("e", false, "Tell the program to encrypt the file")
    decPtr := flag.Bool("d", false, "Tell the program to decrypt the file")
    randKeyPtr := flag.Bool("k", false, "Generates a random key")
    rsaGenPtr := flag.Bool("genRsa", false, "Tell to generate")
    eccGenPtr := flag.Bool("genEcc", false, "Tell to generate")
    sizePtr := flag.Int("size", 16, "Generates a random key of n bytes")

    //Once all flags are declared, call flag.Parse() to execute the command-line parsing.
    flag.Parse()

    var incorrect bool = !*eccGenPtr && !*rsaGenPtr && !*randKeyPtr && *filePtr == "FILE" && *keyPtr == "KEY" && len(flag.Args()) == 0

    if (*encPtr && *decPtr) || (!*eccGenPtr && !*rsaGenPtr && !*encPtr && !*decPtr && !*randKeyPtr) || incorrect {
        flag.Usage()
        os.Exit(0)
    }

    if *namePtr == "NAME" {
        if *encPtr {
            *namePtr = *filePtr + ".enc"
        } else {
            *namePtr = strings.TrimSuffix(*filePtr,".enc")
        }
    }


    var fileName string = *filePtr
    var keyName string = *keyPtr
    var outputFile string = *namePtr

    if *encPtr {

        var key []byte = inout.ReadFile(keyName) //[]byte("example key 1234")
        var file []byte = inout.ReadFile(fileName) //[]byte("exampleplaintext")
        var encrypted []byte = encrypt(file, key)
        
        inout.WriteFile(encrypted, outputFile)
        fmt.Printf("Encrypted %# x\n %# x\n", encrypted[:20], encrypted[len(encrypted)-20:])

        // Mirar si decrypt(encrypted) = file

    } else if *decPtr {

        var file []byte = inout.ReadFile(fileName)
        var key []byte = inout.ReadFile(keyName) 
        var decrypted []byte = decrypt(file, key)

        fmt.Printf("%# x\n", decrypted[len(decrypted)-20:])
        inout.WriteFile(decrypted, outputFile)
    
    } else if *randKeyPtr {
        if *namePtr == "FILE" {
            outputFile = "mySecretKey"
        }
        var newKey []byte = randomKey(*sizePtr)
        inout.WriteFile(newKey,outputFile)
        fmt.Printf("Key generated: %# x\n", newKey)

    } else if *rsaGenPtr {

        RSAkey(*rsaPtr)

    } else if *eccGenPtr {

        ECCkey(*eccPtr)
    }

}
