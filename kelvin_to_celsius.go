package main

import "fmt"

//kelvinToCelsius converte °K in °C
func kelvinToCelsius (k float64) float64{ // Dichiara una funzione che accetta un parametro
	k -= 273.15
	return k //restituisce un risultato
}
func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "°K sono ",celsius,"°C")

}
