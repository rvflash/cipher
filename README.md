# Cipher

[![Build Status](https://img.shields.io/travis/rvflash/cipher.svg)](https://travis-ci.org/rvflash/cipher)
[![Code Coverage](https://img.shields.io/codecov/c/github/rvflash/cipher.svg)](http://codecov.io/github/rvflash/cipher?branch=master)

Ciphers is a just for fun project, created to provide some encryption methods not implemented in the standard crypto/cipher package. 
For now, only the [Caesar cipher](http://github.com/rvflash/cipher/caesar) is available as Go package. 

## Installation

```bash
$ go get github.com/rvflash/cipher
```

## Quick start

If you just need a Go package to play with the Caesar code, see the example bellow :  

```go
import (
	"fmt"
	"github.com/rvflash/cipher/caesar"
)
// ...
r := strings.NewReader("Hello World!")
b, _ := caesar.ROT13.Encrypt(r)
fmt.Printf("%s", b)
// output: Uryyb Jbeyq!
```

As you can see, you can directly use the most famous version of the Caesar cipher, the ROT13 (rotates by 13 places).

> In addiction of the `caesar.ROT13`, you can create your own Caesar cipher with the `caesar.Classic` method.
> It takes as first parameter the number of letter to use as key.
> Finally, the `caesar.New` method uses the same algorithm but don't limit the code to the letters of the alphabet.
> It uses all printable ASCII characters.

If you just want to play with it in your browser, you can start the HTTP server on localhost:8080.

```bash
$ cd $GOPATH/github.com/rvflash/cipher/cmd/cipher
$ go build && ./cipher
```


## Caesar cipher

In cryptography, a Caesar cipher is one of the simplest encryption techniques.
It is a type of substitution cipher in which each letter in the plaintext is replaced by a letter some fixed number of positions down the alphabet.

See the [documentation](https://godoc.org/github.com/rvflash/cipher/caesar) for more details about the interface.


### ROT13

It's the most known version of letter substitution cipher. It replaces a letter with the 13th letter after it.
The variable `caesar.ROT13` gives you a direct access tu use it.