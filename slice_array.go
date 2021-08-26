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
	terrestrial := planets[0:4]
	gasGiants := planets [4:6]
	iceGiants := planets [6:8]
	fmt.Println(terrestrial, gasGiants , iceGiants)

}
