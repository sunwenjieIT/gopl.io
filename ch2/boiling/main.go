// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 29.
//!+

// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
	var test1 Test1 = 100.0
	var test2 Test2
	fmt.Printf("%t\n", test1 == Test1(test2))
	fmt.Print(test1.String())
	fmt.Print(test1)
}

type Test1 float64
type Test2 float64
//!-

func (t Test1) String() string {

	return fmt.Sprintf("%g° C\n'", t)
}