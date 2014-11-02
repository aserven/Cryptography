
package main

import (
    //"encoding/hex"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
    "flag"
    "os"
    "inout"
    "encoding/binary"
    "bytes"
    //"strings"
)

func pkcs7PAD(file []byte, blocksize int) []byte {
    
    var text string = "YELLOW SUBMARINE"
    var text2 string = "YELLOW SUBMA"
    var t []byte = []byte(text)
    var t2 []byte = []byte(text2)
        
    fmt.Println(t)
    fmt.Printf("%# x\n",t2)


    var p int = blocksize - len(file)%blocksize
    
    fmt.Printf("Valor: %d  -- Valor hexa: ",p)
    fmt.Printf("%# x\n",p)
    
    b := make([]byte, 2)
    binary.LittleEndian.PutUint16(b, uint16(p))
    
    b[1] = b[0]

    b = bytes.Repeat(b,p)
    
    fmt.Println(b)
    fmt.Printf("%# x\n",b[len(b)/2:])
    
    return append(file,b[len(b)/2:]...)
    
}

func pkcs7UNPAD(file []byte, blocksize int) []byte {
  
  fmt.Println(file[len(file)-20:])

  b := make([]byte,2)
  
  b[0] = file[len(file)-1]
  b[1] = 0
  
  x := int(binary.LittleEndian.Uint16(b))
  
  fmt.Println(x)
  return file[:len(file)-x]
    
}

func main() {

    filePtr := flag.String("file", "FILE", "File to encrypt/decrypt")
    keyPtr  := flag.String("key", "KEY", "Key to encrypt/decrypt")
    namePtr := flag.String("name", "NAME", "Name of the file to write")
    //This declares numb and fork flags, using a similar approach to the word flag.
    //numbPtr := flag.Int("numb", 42, "an int")
    encPtr := flag.Bool("e", false, "Tell the program to encrypt the file")
    decPtr := flag.Bool("d", false, "Tell the program to decrypt the file")
//It’s also possible to declare an option that uses an existing var declared elsewhere in the program. Note that we need to pass in a pointer to the flag declaration function.
    //var svar string
    //flag.StringVar(&svar, "svar", "bar", "a string var")
//Once all flags are declared, call flag.Parse() to execute the command-line parsing.
    flag.Parse()
//Here we’ll just dump out the parsed options and any trailing positional arguments. Note that we need to dereference the points with e.g. *wordPtr to get the actual option values.
    //fmt.Println("file:", *wordPtr)
    //fmt.Println("key:", *numbPtr)
    //fmt.Println("fork:", *boolPtr)
    //fmt.Println("svar:", svar)
    //fmt.Println("tail:", flag.Args())

    var undeclared bool = *filePtr == "FILE" && *keyPtr == "KEY"

  if (*encPtr && *decPtr || undeclared) {
    flag.Usage();
    os.Exit(0);
  }


  var fileName string = *filePtr;
  var keyName string = *keyPtr;

  key := inout.ReadFile(keyName) //[]byte("example key 1234")
  plaintext := inout.ReadFile(fileName) //[]byte("exampleplaintext")

  plaintext = pkcs7PAD(plaintext,aes.BlockSize)

  fmt.Printf("To encrypt > %# x\n %# x\n", plaintext[:20], plaintext[len(plaintext)-20:])

  // CBC mode works on blocks so plaintexts may need to be padded to the
  // next whole block. For an example of such padding, see
  // https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
  // assume that the plaintext is already of the correct length.
  if len(plaintext)%aes.BlockSize != 0 {
    panic("plaintext is not a multiple of the block size")
  }

  block, err := aes.NewCipher(key)
  if err != nil {
    panic(err)
  }

  // The IV needs to be unique, but not secure. Therefore it's common to
  // include it at the beginning of the ciphertext.
  ciphertext := make([]byte, aes.BlockSize+len(plaintext))
  iv := ciphertext[:aes.BlockSize]
  if _, err := io.ReadFull(rand.Reader, iv); err != nil {
    panic(err)
  }

  fmt.Printf("IV -- %# x\n ",iv)

  mode := cipher.NewCBCEncrypter(block, iv)
  mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

  // It's important to remember that ciphertexts must be authenticated
  // (i.e. by using crypto/hmac) as well as being encrypted in order to
  // be secure.

  inout.WriteFile(ciphertext, "prova"+ ".enc")

  fmt.Printf("Encrypted %# x\n %# x\n", ciphertext[:20], ciphertext[len(ciphertext)-20:])

  //key := []byte("example key 1234")
  //ciphertext2, _ := hex.DecodeString("f363f3ccdcb12bb883abf484ba77d9cd7d32b5baecb3d4b1b3e0e4beffdb3ded")
  
  //ciphertext2 := ciphertext;
  ciphertext2 := inout.ReadFile(fileName+".enc");
  //fmt.Printf("To decrypt %# x\n", ciphertext2)

  iv2 := ciphertext2[:aes.BlockSize]

  block2, err2 := aes.NewCipher(key)
  if err2 != nil {
    panic(err2)
  }

  // The IV needs to be unique, but not secure. Therefore it's common to
  // include it at the beginning of the ciphertext2.
  if len(ciphertext2) < aes.BlockSize {
    panic("ciphertext2 too short")
  }
  //iv := ciphertext2[:aes.BlockSize]
  ciphertext2 = ciphertext2[aes.BlockSize:]

  // CBC mode always works in whole blocks.
  if len(ciphertext2)%aes.BlockSize != 0 {
    panic("ciphertext2 is not a multiple of the block size")
  }

  mode2 := cipher.NewCBCDecrypter(block2, iv2)

  // CryptBlocks can work in-place if the two arguments are the same.
  mode2.CryptBlocks(ciphertext2, ciphertext2)

  // If the original plaintext lengths are not a multiple of the block
  // size, padding would have to be added when encrypting, which would be
  // removed at this point. For an example, see
  // https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
  // critical to note that ciphertexts must be authenticated (i.e. by
  // using crypto/hmac) before being decrypted in order to avoid creating
  // a padding oracle.

  fmt.Printf("Decrypted: %# x\n %# x\n", ciphertext2[:20], ciphertext2[len(ciphertext2)-20:])

  decrypted := pkcs7UNPAD(ciphertext2,aes.BlockSize)
  fmt.Printf("%# x\n", decrypted[len(decrypted)-20:])

  // inout.WriteFile(ciphertext2, strings.TrimSuffix(fileName,".enc"))
  inout.WriteFile(ciphertext2, "prova.jpg")
}
