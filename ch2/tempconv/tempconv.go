// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

//homework code
type Kelvin float64

//米 英尺
type Meter float64
type Foot float64

//公斤/千克 磅
type Kg float64
type Pound float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gT", k) }

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }
func (k Kg) String() string    { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string { return fmt.Sprintf("%glb", p) }

//!-
