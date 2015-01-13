Cryptographic System
==================================

Cryptographic system implemented in Go language for [Cryptography](http://www.fib.upc.edu/fib/estudiar-enginyeria-informatica/assignatures/C.html) subject in [FIB](http://www.fib.upc.edu/fib.html)

General info about Go: http://golang.org/
All libraries used can be visited here: http://golang.org/pkg/

## Install Go ##

Version used: 1.3.3

http://golang.org/doc/install


## Build project ##

Once installed, set the variable GOPATH to the main directory of the project
```bash
export GOPATH="Your directory"/gocode
```

To build the project use

```bash
go install system
```

System is the main directory that contains all rellevant files. The previous comand will create an
executable file in `bin`called `system`. It will also generate another directory called `pkg` 
where the package `inout` will be stored. An overview of the structure:

```
.
├── bin
│   └── system
├── pkg
│   └── darwin_amd64
│       └── inout.a
└── src
    ├── inout
    │   └── inout.go
    └── system
        ├── crypt.go
        ├── keys.go
        ├── main.go
        ├── message.go
        └── signature.go
```
* __`crypt.go`__: Contains all functions for encrypt and decrypt a file
* __`keys.go`__: Functions to generate an RSA key and EC key
* __`signature.go`__: Functions to sign and verify files with EC and RSA keys
* __`message.go`__: Functions to send and receive a message
* __`main.go`__: Main program
* __`inout.go`__: Functions to read and write files


### Usage ###
```
--------------------------------------------------------------------------------
USAGE: system -action < ACTION >
--------------------------------------------------------------------------------
     encrypt -file FILE -key KEY [-name OUTPUT]
     decrypt -file FILE -key KEY [-name OUTPUT]
     rsa     [-size INTEGER]
     ec      [-name <256 | 384 | 521>]
     sign    -file FILE -key ECKEY [-name OUTPUT]
     verify  -file FILE -key ECKEY  -sign SIGNATURE
     send    -file FILE -key RSAKEY -sign ECKEY [-name OUTPUT]
     receive -file FILE -key RSAKEY -sign ECKEY [-name OUTPUT] [-signName SIGNATURE]

  * encrypt: Encrypts a file with a key usgin AES operating with CBC
  * decrypt: Decrypts a file with a key usgin AES operating with CBC
  * rsa:     Generates an RSA key in PEM format (public and private files). DEFAULT VALUE: 2048
  * ec:      Generates an EC key in PEM format (public and private files). DEFAULT CURVE: P256
  * sign:    Signs a file with the given key. Has to be EC
  * verify:  Verifies if the file is signed with the signature given
  * send:    Signs the message and encrypts it with AES including the key encrypted with RSA key
  * receive: Decrypts the file and writes the message and the signature

  -file FILE: Name of the file to use
  -key  KEY:  Name of the key to use
  -sign SIGN: Name of the key to sign/verify or signature
  -name NAME: Name of the file to be written or curve to be used
  -size SIZE: Number of bits to generate the RSA key
  -signName NAME: Name of the file to be written as signature
--------------------------------------------------------------------------------
```

### Examples of calls ###

```
   system -action encrypt ../test/bits.jpg -key ../test/SecretKey 
   system -action encrypt -file ../test/bits.jpg -key ../test/SecretKey 
   system -action decrypt -file result.out -key ../test/SecretKey 
   system -action rsa
   system -action ec
   system -action sign -file ../test/bits.jpg -key privateEC.pem 
   system -action verify -file result.out -key publicEC.pem 
   system -action verify -file ../test/bits.jpg -sign result.out -key publicEC.pem 
   system -action send -file ../test/bits.jpg -key privateRSA.pem -sign privateEC.pem 
   system -action send -file ../test/bits.jpg -key publicRSA.pem -sign privateEC.pem -name bits.jpg.enc
   system -action receive -file bits.jpg.enc -key privateRSA.pem -sign publicEC.pem -name bitsOK.jpg -signName bits.jpg.signature
```

## Details ##
### 1. Encrypting / Decrypting ###
Encryption and decryption uses the AES algorithm with blocks of 128 bits operating with
cypher block chaining (CBC) with the given key (usually 16 bytes). An IV vector is added 
as header and before encryption and after decryption uses padding PKCS #7

### 2. Generating RSA and ECC keys ###

#### 2.1 RSA ####
Given an integer n, generates an RSA key with n bits in PEM format. Default value is 2048 bits (recommended).
Outputs always two files:
   - publicRSA.pem
   - privateRSA.pem

If you want to check your generated key you can run in the terminal:

```
openssl rsa -in publicRSA.pem -pubin -tex
openssl rsa -in privateRSA.pem -text
```

#### 2.2 EC ####
Given a name of a curve (only available to input 256, 384 and 521) generates that 
curve (prime256v1,prime384v1,prime521v1). By default generates prime256v1.

If you want to check your generated key you can run in the terminal:
```
openssl ec -in publicEC.pem -pubin -text
openssl ec -in privateEC.pem -text
```

To list all the curves available in openssl use
```openssl ecparam -list_curves```

### 3. Sign / Verify ###
The system can sign with RSA and EC keys. The functions in the main program are with EC. 
All signatures and verifications use SHA256 hash.

To sign a file is needed to provide its name and the private key. To verify it needs the public
key and also the signature. With EC keys generates two big ints, each one has 256 bits, so the 
size of the signature is 64 bytes.

### 4. Send / Receive Message ###
The function that sends a message is provided with the file, a public RSA key (2048 bits) and a private EC
key (P256) for signing purposes. First, signs the file with the private EC key, then appends the signature to
the file and encrypts all with AES with a previously generated random key of 16 bytes. The key used is encrypted 
with the public RSA key and added as header.

The function that receives the message does the inverse process. Given the encrypted file, a private RSA key and
a public EC key, first takes the first 2048 bits and decrypts them with the private RSA key. Now decrypts the other
part with the key and separates the message from the signature (Last 64 bytes).




