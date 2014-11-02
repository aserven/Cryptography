
package main

import (
    "os"
    "flag"
    "strings"
)



func main() {

    filePtr := flag.String("file", "FILE", "File to encrypt/decrypt")
    keyPtr  := flag.String("key", "KEY", "Key to encrypt/decrypt")
    namePtr := flag.String("name", "NAME", "Name of the file to write")
    encPtr := flag.Bool("e", false, "Tell the program to encrypt the file")
    decPtr := flag.Bool("d", false, "Tell the program to decrypt the file")

    //Once all flags are declared, call flag.Parse() to execute the command-line parsing.
    flag.Parse()

    var incorrect bool = *filePtr == "FILE" && *keyPtr == "KEY" && len(flag.Args()) == 0

    if (*encPtr && *decPtr) || (!*encPtr && !*decPtr) || incorrect {
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
        encrypt(fileName, keyName, outputFile)
    } else if (*decPtr) {
        decrypt(fileName, keyName, outputFile)
    }

}
