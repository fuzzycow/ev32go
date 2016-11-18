// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")


func main() {
	log.Printf("Starting")
	flag.Parse()
	log.Fatal(http.ListenAndServe(*addr, http.FileServer(http.Dir("."))))


}
