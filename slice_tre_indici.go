package main

import "fmt"

func main() {
	planets := [8] string{
		"Mercurio",
		"Venere",
		"Terra",
		"Marte",
		"Giove",
		"Saturno",
		"Urano",
		"Nettuno",
	}
	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Cerere")
	fmt.Println(planets)
	terrestrial = planets [0:4]
	worlds = append(terrestrial,"Cerere")
	fmt.Println(planets)
	fmt.Print(worlds)
}
