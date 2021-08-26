package main

import "fmt"

func main() {
	var planets [8] string
	planets[0] = "Mercurio"
	planets[1] = "Venere"
	planets[2] = "Terra"
	earth := planets[2]
	fmt.Println(earth)
	fmt.Println(len(planets))
}
