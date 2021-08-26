package main

import (
	"fmt"
	"math/rand"
)

/*
	La distanza fra la Terra e Marte varia a seconda del fatto che si trovino dalla
	stessa parte o dalla parte opposta rispetto al Sole.
	Scrivete un programma che generi una distanza causale compresa fra 56 e 401 milioni di Km
 */
func main() {
var distanza = rand.Intn((350) + 51) * 1000000
fmt.Println("La distanza tra la Terra e Marte è di ", distanza, "Km")
}
