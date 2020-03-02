package main

import (
	"fmt"
)

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Print(k, "°Kは、", c, "°Cです.\n")

	var k2 kelvin = 294.0
	c2 := k2.celsius()
	fmt.Print(k2, "°Kは、", c2, "°Cです.\n")
}
