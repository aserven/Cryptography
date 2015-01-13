
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
    signaturePtr := flag.String("signName", "Signed", "Name of the file to write")
    rsaPtr := flag.Int("rsa", 2048, "Generate RSA key of n bits")
    eccPtr := flag.String("ECC", "256", "Generate ECC key called P256, P384 or P521")
    encPtr := flag.Bool("e", false, "Tell the program to encrypt the file")
    decPtr := flag.Bool("d", false, "Tell the program to decrypt the file")
    randKeyPtr := flag.Bool("k", false, "Generates a random key")
    rsaGenPtr := flag.Bool("genRsa", false, "Tell to generate")
    eccGenPtr := flag.Bool("genEcc", false, "Tell to generate")
    signPtr := flag.Bool("sign", false, "Tell to sign file")
    verifyPtr := flag.Bool("verify", false, "Tell to verify file")
    sizePtr := flag.Int("size", 16, "Generates a random key of n bytes")
    test := flag.Bool("test", false, "Test program")

    //Once all flags are declared, call flag.Parse() to execute the command-line parsing.
    flag.Parse()

    var incorrect bool = !*verifyPtr && !*signPtr && !*eccGenPtr && !*rsaGenPtr && !*randKeyPtr && *filePtr == "FILE" && *keyPtr == "KEY" && len(flag.Args()) == 0

    if (!*test) && ((*encPtr && *decPtr) || (!*verifyPtr && !*signPtr && !*eccGenPtr && !*rsaGenPtr && !*encPtr && !*decPtr && !*randKeyPtr) || incorrect) {
        // FER EL PROPI
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

    var signName string = *signaturePtr
    var file []byte = inout.ReadFile(fileName)
    var key []byte = inout.ReadFile(keyName) 
    var sign []byte = inout.ReadFile(signName)

    //var result []byte = sendMessage(file,key,sign)

    //inout.WriteFile(result,"FileENC.enc")
    //fmt.Printf("RESULT: %# x\n%# x", result[:256],result[len(result)-64:])


    Mess, Firm := receiveMessage(file,key,sign)

    inout.WriteFile(Mess,"MissatgeENC.png")
    inout.WriteFile(Firm,"FirmaENC.signature")


    //QUAN FA ENCODE FER DECODE I COMPROVAR QUE ESTA CORRECTE
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

    }else if *signPtr {
        var file []byte = inout.ReadFile(fileName)
        var key []byte = inout.ReadFile(keyName) 
        //var signed []byte = sign(file,key)
        var signed2 []byte = signECDSA(file,key)

        //nameVec := strings.Split(fileName, ".")
        //outFile := strings.Join(nameVec[0:len(nameVec)-1],"") + "SIGNED." + nameVec[len(nameVec)-1]
        //inout.WriteFile(signed,"testRSA.signature")
        inout.WriteFile(signed2,"testECDS2.signature")
        
    }else if *verifyPtr {
        var signName string = *signaturePtr
        var file []byte = inout.ReadFile(fileName)
        var key []byte = inout.ReadFile(keyName) 
        var sign []byte = inout.ReadFile(signName)
        
        //var ret bool = verify(file,sign,key)
        var ret bool = verifyECDSA(file,sign,key)
        if ret {
            fmt.Println("True")
        } else {
            fmt.Println("False")
        }
    }

}
