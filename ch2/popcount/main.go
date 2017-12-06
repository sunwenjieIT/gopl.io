// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

var pc2 [65536]byte

func init() {
	for i := range pc {
		//pc[i] = pc[i/2] + byte(i&1)
		pc[i] = pc[i>>1] + byte(i&1)
	}

	for i := range pc2 {
		//pc2[i] = pc2[i/2] + byte(i&1)

		pc2[i] = pc2[i>>1] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount2(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//!-

// homework 	重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。 （
// 11.4节将展示如何系统地比较两个不同实现的性能。）

func PopCount(x uint64) int {

	var i uint8 = 0
	var sum byte

	for i = 0; i < 8; i++ {
		tmp := pc[byte(x>>(i*8))]
		fmt.Printf("%d\n", tmp)
		sum += tmp
	}

	return int(sum)
}

//查表算法 8位改16位
func PopCount3(x uint64) int {

	return int(pc2[uint16(x>>(0*16))] +
		pc2[uint16(x>>(1*16))] +
		pc2[uint16(x>>(2*16))] +
		pc2[uint16(x>>(3*16))])

}