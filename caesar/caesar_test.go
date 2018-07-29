// Copyright (c) 2018 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package caesar_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/rvflash/cipher/caesar"
)

func ExampleClassic() {
	r := strings.NewReader("Hello World!")
	b, _ := caesar.ROT13.Encrypt(r)
	fmt.Printf("%s", b)
	// output: Uryyb Jbeyq!
}

const (
	poemPath          = "./testdata/pierre-ronsard.txt"
	poemCipherPath    = "./testdata/pierre-ronsard.cipher.txt"
	poemNewCipherPath = "./testdata/pierre-ronsard.new.cipher.txt"
)

var data = map[string]string{
	"hi":           "Hello World!",
	"hi_13":        "Uryyb Jbeyq!",
	"hi_13_p":      "Uryy|-d| yq.",
	"AZ":           "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"AZ_23":        "XYZABCDEFGHIJKLMNOPQRSTUVW",
	"AZ_23_p":      "XYZ[\\]^_`abcdefghijklmnopq",
	"TEXT":         "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG",
	"TEXT_23":      "QEB NRFZH YOLTK CLU GRJMP LSBO QEB IXWV ALD",
	"TEXT_23_p":    "k_\\7hl`Zb7Yifne7]fo7aldgj7fm\\i7k_\\7cXqp7[f^",
	"Unicode":      "Hello World :) 😀!",
	"Unicode_13":   "Uryyb Jbeyq :) 😀!",
	"Unicode_13_p": "Uryy|-d| yq-G6-😀.",
}

// R is a short cut to the strings.NewReader method.
var R = strings.NewReader

func TestCaesar(t *testing.T) {
	var dt = []struct {
		c       *caesar.Caesar
		in, out io.Reader
		decrypt bool
	}{
		// Classic
		{c: caesar.ROT13},
		// 1
		{c: caesar.ROT13, in: R("")},
		{c: caesar.ROT13, in: R(data["hi"]), out: R(data["hi_13"])},
		{c: caesar.ROT13, in: R(data["hi_13"]), out: R(data["hi"]), decrypt: true},
		{c: caesar.ROT13, in: R(data["Unicode"]), out: R(data["Unicode_13"])},
		{c: caesar.ROT13, in: R(data["Unicode_13"]), out: R(data["Unicode"]), decrypt: true},
		{c: caesar.Classic(23), in: R(data["AZ"]), out: R(data["AZ_23"])},
		{c: caesar.Classic(23), in: R(data["AZ_23"]), out: R(data["AZ"]), decrypt: true},
		{c: caesar.Classic(23), in: R(data["TEXT"]), out: R(data["TEXT_23"])},
		{c: caesar.Classic(23), in: R(data["TEXT_23"]), out: R(data["TEXT"]), decrypt: true},
		// 10
		{c: caesar.Classic(3).Reverse(), in: R(data["TEXT"]), out: R(data["TEXT_23"])},
		{c: caesar.Classic(3).Reverse(), in: R(data["TEXT_23"]), out: R(data["TEXT"]), decrypt: true},
		// New (with all printable ASCII)
		{c: caesar.New(13), in: R(data["hi"]), out: R(data["hi_13_p"])},
		{c: caesar.New(13), in: R(data["hi_13_p"]), out: R(data["hi"]), decrypt: true},
		{c: caesar.New(13), in: R(data["Unicode"]), out: R(data["Unicode_13_p"])},
		{c: caesar.New(13), in: R(data["Unicode_13_p"]), out: R(data["Unicode"]), decrypt: true},
		{c: caesar.New(23), in: R(data["AZ"]), out: R(data["AZ_23_p"])},
		{c: caesar.New(23), in: R(data["AZ_23_p"]), out: R(data["AZ"]), decrypt: true},
		{c: caesar.New(23), in: R(data["TEXT"]), out: R(data["TEXT_23_p"])},
		{c: caesar.New(23), in: R(data["TEXT_23_p"]), out: R(data["TEXT"]), decrypt: true},
		// 20
		{c: caesar.Classic(13), in: openFile(poemPath), out: openFile(poemCipherPath)},
		{c: caesar.New(13), in: openFile(poemPath), out: openFile(poemNewCipherPath)},
	}
	var (
		b   []byte
		err error
	)
	for i, tt := range dt {
		if tt.decrypt {
			b, err = tt.c.Decrypt(tt.in)
		} else {
			b, err = tt.c.Encrypt(tt.in)
		}
		if err != nil {
			t.Fatalf("%d. unexpected error: got=%q", i, err)
		}
		if s := read(tt.out); string(b) != s {
			t.Errorf("%d. mismatch content: got=%q exp=%q", i, b, s)
		}
	}
}

func openFile(name string) io.Reader {
	f, err := os.Open(name)
	if err != nil {
		return nil
	}
	return f
}

func read(r io.Reader) string {
	if r == nil {
		return ""
	}
	in := new(bytes.Buffer)
	if _, err := in.ReadFrom(r); err != nil {
		return ""
	}
	return in.String()
}

func TestCaesar_Rune(t *testing.T) {
	var dt = []struct {
		c   *caesar.Caesar
		in  rune
		out string
	}{
		{c: caesar.ROT13, in: 'a', out: "n"},
		{c: &caesar.Caesar{}, in: 'a', out: "\x00"},
	}
	for i, tt := range dt {
		if out := tt.c.Rune(tt.in, false); out.String() != tt.out {
			t.Errorf("%d. content mismatch: got=%s exp=%s", i, out, tt.out)
		}
	}
}
