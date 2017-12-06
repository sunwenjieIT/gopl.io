// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//!-

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

//米 英尺 转换
func MToFt(m Meter) Foot {
	return Foot(3.2808399 * m)
}

func FtToM(f Foot) Meter {
	return Meter(f / 3.2808399)
}

//千克 磅 转换
func KgToLb(kg Kg) Pound {
	return Pound(kg * 2.2046226)
}

func LbToKg(pound Pound) Kg {
	return Kg(pound / 2.2046226)
}