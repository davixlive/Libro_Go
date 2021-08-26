package main

import "fmt"

func kelvinToCelsius (k float64) float64{ // Dichiara una funzione che accetta un parametro
	k -= 273.15
	return k //restituisce un risultato
}
func celsiusToFahrenheit (c float64) float64{
	f := (c * 9.0 / 5.0 ) + 32.0
	return f
}
func main() {
	k := 0.0
	f := celsiusToFahrenheit(kelvinToCelsius(k))
	fmt.Print(k,"°K equivalgono a ",f,"° F")
}
