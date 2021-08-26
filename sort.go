package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := [] string{
		"Mercurio",
		"Venere",
		"Terra",
		"Marte",
		"Giove",
		"Saturno",
		"Urano",
		"Nettuno",
	}
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
	sort.Strings(planets)
	fmt.Println(planets)
}
