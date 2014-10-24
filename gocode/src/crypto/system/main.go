/*package main

import (
    "fmt"
    "crypto/crypt"
    "crypto/string"
    "crypto/keys"
    "crypto/signature"
    "crypto/read"
)

func main() {


    var file []byte = read.ReadFile("test.txt")

    var key string = []byte("")

    var fileEn = crypt.encrypt(file,key)

    fmt.Println(string.Reverse("Hello, new gopher!"))
}*/


package main
 
import (
  "crypto/aes"
  "crypto/cipher"
  "fmt"
)
 
func main() {
  input := []byte("this is a test")
  iv := []byte("532b6195636c6127")[:aes.BlockSize]
  key := []byte("532b6195636c61279a010000")
 
  fmt.Println("Input:     ", input)
  fmt.Println("Key:       ", key)
  fmt.Println("IV:        ", iv)
 
  encrypted := make([]byte, len(input))
  EncryptAES(encrypted, input, key, iv)
 
  fmt.Println("Output:    ", encrypted)

  var decrypted
  DecryptAES(decrypted, encrypted, key, iv)

  fmt.Println
}
 
func EncryptAES(dst, src, key, iv []byte) error {
  aesBlockEncryptor, err := aes.NewCipher([]byte(key))
  if err != nil {
    return err
  }
  aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncryptor, iv)
  aesEncrypter.XORKeyStream(dst, src)
  return nil
}
 
func DecryptAES(dst, src, key, iv []byte) error {
  aesBlockEncryptor, err := aes.NewCipher([]byte(key))
  if err != nil {
    return err
  }
  aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncryptor, iv)
  aesEncrypter.XORKeyStream(dst, src)
  return nil
}