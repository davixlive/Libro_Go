package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var moneta5, moneta10, moneta20 int
	for salvadanaio:= 0.0;salvadanaio<20;{
		switch monetacasuale:=rand.Intn(3); monetacasuale {
		case 0:
			moneta5++
		case 1:
			moneta10++
		case 2:
			moneta20++
		}
		salvadanaio= float64(moneta5)*0.05 + float64(moneta10)*0.1 + float64(moneta20)*0.2
		fmt.Printf("Saldo %4.2f\n",salvadanaio)
	}

}
