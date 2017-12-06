// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

//修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "first fetch: %v\n", err)
			os.Exit(1)
		}
		_, newerr := io.Copy(os.Stdout, resp.Body)

		//b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if newerr != nil {
			fmt.Fprintf(os.Stderr, "resp.Status: %s.\n", resp.Status)
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
		//fmt.Printf("%v", resp.Status)
	}
}

//!-
