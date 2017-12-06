// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
	"bufio"
	"flag"
)

var cfType = flag.String("type", "length", "需要执行的换算类型, 默认为 length: 长度换算, 其他可选项: temperature 温度 mass 质量 , ")

func doConvert(t float64)  {
	switch *cfType {

	case "length":
		m := tempconv.Meter(t)
		ft := tempconv.Foot(t)
		fmt.Printf("%s = %s, %s = %s\n",
			m, tempconv.MToFt(m), ft, tempconv.FtToM(ft))
	case "temperature":
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	case "mass":
		kg := tempconv.Kg(t)
		lb := tempconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n",
			kg, tempconv.KgToLb(kg), lb, tempconv.LbToKg(lb))
	}
}

//通用的单位转换程序, 支持 米/英尺 千克/磅 摄氏度/华氏度
func main() {
	//flag.Parse()


	for index, arg := range flag.Args() {
		fmt.Printf("index: %d, value: %s.\n", index, arg)
	}
	if len(flag.Args()) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			if input.Text() == "end" {
				break
			}
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v. please retype. \n", err)
				os.Exit(1)
			}
			doConvert(t)
		}
		return
	}

	fmt.Printf("command with args.\n")
	for _, arg := range flag.Args() {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		doConvert(t)
	}

}

//!-
