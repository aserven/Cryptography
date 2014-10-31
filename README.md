Sistema Criptogràfic
==================================

Implementació d'un sistema criptogràfic en el llenguatge Go

http://golang.org/

Instal·lació: (Versió usada: 1.3.3)

[Link](http://golang.org/doc/install)

GOPATH="Your path where you want the code to be stored"


tree source code:

```
.
├── crypto
│   ├── crypt
│   │   └── crypt.go
│   ├── keys
│   │   └── keys.go
│   ├── signature
│   │   └── signature.go
│   ├── string
│   │   ├── string.go
│   │   └── string_test.go
│   └── system
│       └── main.go
├── read
│   └── read.go
└── system

```

Encrypting Decrypting

AES 128 block  operating with CBC chiper block chaining  with key 16 bytes


padding pcks7

MODULES:

* __crypt__ : encrypt and decrypt functionalities
* __keys__ : generators of keys RSA and eliptic
* __signature__ : Methods to sign and authenticate
* __main__ : Main program 

- _system_ : 
- _inout_ : Package to read and write files


