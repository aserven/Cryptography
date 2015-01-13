Cryptographic System
==================================

Cryptographic system implemented in Go language for [Cryptography](http://www.fib.upc.edu/fib/estudiar-enginyeria-informatica/assignatures/C.html) subject in [FIB](http://www.fib.upc.edu/fib.html)

General info about Go: http://golang.org/

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

System 

 - install places the binary in `bin`
 - builds places the binary as a package in `pkg`

tree structure of the go project:

```
.
├── bin
│   └── system
├── pkg
│   └── darwin_amd64
│       ├── crypto
│       │   ├── keys.a
│       │   ├── signature.a
│       │   └── string.a
│       └── inout.a
│  
└── src
    ├── inout
    │   └── inout.go
    └── system
        ├── crypt.go
        ├── keys.go
        ├── main.go
        └── signature.go
```

USAGE:
Usage of system:
  - d=false: Tell the program to decrypt the file
  - e=false: Tell the program to encrypt the file
  - file="FILE": File to encrypt/decrypt
  - key="KEY": Key to encrypt/decrypt
  - name="NAME": Name of the file to write


### 1. Encrypting Decrypting ###

AES 128 block  operating with CBC chiper block chaining  with key 16 bytes

iv  vector added as header

padding pcks7

### 2. Generating RSA and ECC keys ###

options: [n] number of bytes
 
* public key: {m, e}
* private key: {m,e,d,p,q,dp,dq,Q}

_(In PEM format: private.pem, public.pem)_
ASN.1, DER format --> then generate PEM file

```
openssl rsa -in private.pem -text
openssl rsa -in public.pem -pubin -text
```

options: [curve] name of the curve

curvas secp256, secp384 y secp521
curvas secp256, secp384 y secp521
generate 256, 384 and 521 curves

```openssl ecparam -list_curves```

```
openssl ec -in private.pem -text
openssl ec -in public.pem -pubin -text
```

### 3. Sign Verify ###

* RSA (signature 256 bytes with 256 bits)
    - SignPSS 
    - VerifyPSS

* ECDSA with P256 (signature occupies 2 x 256 bytes, generates two big ints)
    - Sign
    - Verify

**to sign:** file key - returns file.signature
**to verify:** file key signature - returns True | False 

with SHA256 hash

### 4. Send / Receive ###

should use ecdsa with DH

## MODULES: ###

* __crypt__ : encrypt and decrypt functionalities
* __keys__ : generators of keys RSA and eliptic
* __signature__ : Methods to sign and authenticate
* __main__ : Main program 


- _system_ : All cryptography modules inside
- _inout_ : Package to read and write files



EXAMPLES

    bin/system -action encrypt ../test/bits.jpg -key ../test/SecretKey 
    bin/system -action encrypt -file ../test/bits.jpg -key ../test/SecretKey 
    bin/system -action decrypt -file result.out -key ../test/SecretKey 
    bin/system -action rsa
    bin/system -action ec
    bin/system -action sign -file ../test/bits.jpg -key privateEC.pem 
    bin/system -action verify -file result.out -key publicEC.pem 
    bin/system -action verify -file ../test/bits.jpg -sign result.out -key publicEC.pem 
    bin/system -action send -file ../test/bits.jpg -key privateRSA.pem -sign privateEC.pem 
    bin/system -action send -file ../test/bits.jpg -key publicRSA.pem -sign privateEC.pem -name bits.jpg.enc
    bin/system -action receive -file bits.jpg.enc -key privateRSA.pem -sign publicEC.pem -name bitsOK.jpg -signName bits.jpg.signature

