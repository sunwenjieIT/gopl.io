package main

import (
	"os"
	"strconv"
	"fmt"
	"gopl.io/ch1/test"
)

//欧几里德GCD算法验证
func main() {

	/*
	for i, v := range os.Args {

		fmt.Printf("index: %d, value: %s.\n", i, v)
	}
	*/

	test.PrintArgs()

	if len(os.Args) > 2 {

		if x, err := strconv.Atoi(os.Args[1]); err == nil {
			if y, err := strconv.Atoi(os.Args[2]); err == nil {

				result := gcd(x, y)

				fmt.Printf("gcd result: %d.\n", result)
			}
		}
	}

}

func gcd(x, y int) int {

	for y != 0 {
		x, y = y, x%y
	}

	return x
}
