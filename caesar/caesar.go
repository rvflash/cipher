// Copyright (c) 2018 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

// Package caesar provides interface to encrypt and decrypt Caesar cipher.
package caesar

import (
	"bytes"
	"fmt"
	"io"
)

var (
	// Classical bounds of the cipher.
	alphabet = []int{'a', 'z', 'A', 'Z'}
	// Printable ASCII as bounds.
	printable = []int{' ', '~'}
)

// ROT13 rotates by 13 places. It's the most famous Caesar cipher.
var ROT13 = &Caesar{shift: 13, bounds: alphabet}

// Caesar represents a Caesar cipher.
// For more information on the Caesar's code:
// https://en.wikipedia.org/wiki/Caesar_cipher
type Caesar struct {
	shift  int
	left   bool
	bounds []int
}

// Classic returns a classic Caesar cipher, only based on the alphabet.
// So, only the letter a to z or A to Z are substitute.
func Classic(key int) *Caesar {
	return &Caesar{shift: key, bounds: alphabet}
}

// New returns a Caesar cipher using all the ASCII printable characters.
// It extends the number of characters classically used
// See bellow the complete list:
// > https://en.wikipedia.org/wiki/ASCII#Printable_characters
func New(key int) *Caesar {
	return &Caesar{shift: key, bounds: printable}
}

// Encrypt uses the current cipher to encrypt the given stream.
func (c *Caesar) Encrypt(r io.Reader) ([]byte, error) {
	return c.write(r, false)
}

// Decrypt uses the current cipher to decrypt the given stream.
func (c *Caesar) Decrypt(r io.Reader) ([]byte, error) {
	return c.write(r, true)
}

func (c *Caesar) write(r io.Reader, reverse bool) (buf []byte, err error) {
	if r == nil {
		return
	}
	in := new(bytes.Buffer)
	if _, err = in.ReadFrom(r); err != nil {
		return
	}
	out := new(bytes.Buffer)
	for _, r := range in.String() {
		if _, err = out.WriteString(c.Rune(r, reverse).String()); err != nil {
			return
		}
	}
	return out.Bytes(), nil
}

// Rune applies the code on the given rune.
// The return implements the fmt.Stringer interface.
func (c *Caesar) Rune(r rune, reverse bool) fmt.Stringer {
	return &code{cipher: c, reverse: reverse, rune: int(r)}
}

// Reverse reverses the direction of the substitution.
// If actually defined to right, the substitution will go the right after it.
func (c *Caesar) Reverse() *Caesar {
	c.left = !c.left
	return c
}

type code struct {
	cipher  *Caesar
	reverse bool
	rune    int
}

// Returns three parameters:
// The first, a boolean returns true if the rune keeps inside the bounds.
// Then the minimum and maximum of the bounded range.
func (c code) bounded() (ok bool, min, max int) {
	if ok = len(c.cipher.bounds) == 0; ok {
		// Deals with empty struct.
		return
	}
	for i := 1; i < len(c.cipher.bounds); i = i + 2 {
		if ok = c.cipher.bounds[i-1] <= c.rune && c.rune <= c.cipher.bounds[i]; ok {
			min, max = c.cipher.bounds[i-1], c.cipher.bounds[i]
			return
		}
	}
	return
}

// Returns true if the substitution goes to the left.
func (c code) leftShifted() bool {
	if c.reverse {
		return !c.cipher.left
	}
	return c.cipher.left
}

// String implements the fmt.Stinger interface.
func (c code) String() string {
	ok, min, max := c.bounded()
	if !ok {
		// Do not care to this rune.
		return string(c.rune)
	}
	// Direction.
	diff := max - min + 1
	shift := c.cipher.shift
	if c.leftShifted() {
		shift = diff - shift%diff
	}
	return string((c.rune-min+shift)%diff + min)
}
