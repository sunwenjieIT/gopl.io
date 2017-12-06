package main

import (
	"os"
	"strconv"
	"gopl.io/ch2/popcount"
	"fmt"
)

func main() {

	//os.Args
	if len(os.Args) == 1 {
		r := popcount.PopCount3(0x1234567890ABCDEF)
		fmt.Printf("result: %d.\n", r)
		os.Exit(1)
	}

	for _, v := range os.Args[1:] {

		x, err :=strconv.ParseUint(v, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		r := popcount.PopCount(x)
		fmt.Printf("%s has %d 1.\n", v, r)
	}

}
