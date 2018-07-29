// Copyright (c) 2018 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/rvflash/cipher/caesar"
)

func caesarHandler(w http.ResponseWriter, r *http.Request) {
	// Parses the query string.
	q := r.URL.Query()
	s, ok := q["s"]
	if !ok {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Encrypts it!
	res, err := caesar.ROT13.Encrypt(strings.NewReader(s[0]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Allows access to the API with JS calls.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "./static/index.html")
}

func main() {
	// Environment.
	port := flag.Int("port", 4433, "service port")
	flag.Parse()

	// Caesar cipher handler.
	http.HandleFunc("/caesar", caesarHandler)
	// Landing page.
	http.HandleFunc("/", defaultHandler)

	// Start the server.
	addr := ":" + strconv.Itoa(*port)
	fmt.Printf("Listening on localhost%s\n", addr)
	err := http.ListenAndServeTLS(addr, "./testdata/server.pem", "./testdata/server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
