package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {

	PrintArgs()
	/*
	for k, v := range os.Args {
		fmt.Printf("index: %d, value: %s.\n", k, v)
	}
	*/

	if len(os.Args) > 1 {

		if x, err := strconv.Atoi(os.Args[1]); err == nil {
			x = fibonacci(x)
			fmt.Printf("fibonacci result: %d.\n", x)
		}
	}

}

func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func PrintArgs() {
	for k, v := range os.Args {
		fmt.Printf("index: %d, value: %s.\n", k, v)
	}
}
