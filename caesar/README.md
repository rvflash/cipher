# Caesar cipher

[![GoDoc](https://godoc.org/github.com/rvflash/cipher/caesar?status.svg)](https://godoc.org/github.com/rvflash/cipher/caesar)
[![Go Report Card](https://goreportcard.com/badge/github.com/rvflash/cipher/caesar)](https://goreportcard.com/report/github.com/rvflash/cipher/caesar)

In cryptography, a Caesar cipher is one of the simplest encryption techniques.
It is a type of substitution cipher in which each letter in the plaintext is replaced by a letter some fixed number of positions down the alphabet.

2 methods are available:
* `caesar.Classic` implements the well known Caesar code. It takes as argument the key to use to shit letter between `a` to `z` or `A` to `Z`. 
* `caesar.New` implements the same algorithm but do not limit it to the alphabet. It uses all printable ASCII characters.

The `caesar.ROT13` variable allows to quickly use the most famous version of the Caesar code.

See the [documentation](https://godoc.org/github.com/rvflash/cipher/caesar) for more details about the interface.


## Quick start

```
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