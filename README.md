Sistema Criptogràfic
==================================

Implementació d'un sistema criptogràfic en el llenguatge Go

http://golang.org/

Instal·lació: (Versió usada: 1.3.3)

[Download here](http://golang.org/doc/install)

GOPATH="Your path where you want the code to be stored"


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


Encrypting Decrypting

AES 128 block  operating with CBC chiper block chaining  with key 16 bytes

iv  vector added as header

padding pcks7

MODULES:

* __crypt__ : encrypt and decrypt functionalities
* __keys__ : generators of keys RSA and eliptic
* __signature__ : Methods to sign and authenticate
* __main__ : Main program 


- _system_ : All cryptography modules inside
- _inout_ : Package to read and write files


Generating RSA key options: [n] number of bytes

* public key: {m, e}
* private key: {m,e,d,p,q,dp,dq,Q}

_(In PEM format)_


