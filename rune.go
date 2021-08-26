package main

import "fmt"

/*
	type byte = uint8
	type rune = int32
 */
func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang )
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang ) //visualizza i caratteri al posto del valore numerico
}
