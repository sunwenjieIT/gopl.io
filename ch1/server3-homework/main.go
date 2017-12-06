// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"log"
	"net/http"
	"gopl.io/ch1/test"
)


//修改lissajour服务, 从URL读取变量. cycles默认5, 通过http from传入
func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		test.Lissajous(w, r)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}



//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		//fmt.Printf("Header[%q] = %q\n", k, v)
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

//!-handler
