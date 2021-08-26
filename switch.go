package main

import "fmt"

func main() {
	var room = "lago"
	switch {
	case room == "caverna":
		fmt.Println("Ti trovi in una caverna buia.")
	case room == "lago":
		fmt.Println("Il ghiaccio sembra reggere.")
		fallthrough // Permette di passare al case successivo
	case room == "sottacqua":
		fmt.Println("L'acqua è ghiacciata.")
	}
}
