package main

import "fmt"

func main() {

	const distanza =  56000000 //km
	const giorni = 28
	var velocita = distanza / 28 * 24
	fmt.Println("Per raggiungere Malacandra in 28 giorni, dovremmo viaggiare ad una velocità di ", velocita, "km/h")
}
