package main

import "fmt"

type celsius float64
type kelvin float64
func kelvinToCelsius (k kelvin)  celsius{
	return celsius(k - 273.15)
}
func (k kelvin) celsius() celsius{
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	var c celsius
	c = kelvinToCelsius(k)
	fmt.Println(c)
	c = k.celsius()
	fmt.Print(c)
}
