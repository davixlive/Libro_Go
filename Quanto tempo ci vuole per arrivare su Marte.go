package main

import "fmt"

func main() {
	const velocitaluce = 299792 // Km/s
	var distanza = 56000000 //km
	fmt.Println(distanza/velocitaluce, "secondi")
	distanza = 401000000
	fmt.Println(distanza/velocitaluce, "secondi")

}
