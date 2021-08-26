package main

import "fmt"

func main() {
	planets := [8] string {
		"Mercurio",
		"Venere",
		"Terra",
		"Marte",
		"Giove",
		"Saturno",
		"Urano",
		"Nettuno",
	}
	for _,planet := range planets{
		fmt.Println(planet)
	}
}
