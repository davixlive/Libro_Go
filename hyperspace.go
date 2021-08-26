package main

import (
	"fmt"
	"strings"
)
//hyperspace rimuove lo spazio che circonda i vari mondi
func hyperspace (worlds []string) { //questo argomento è una slice non un array
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
func main() {
	planets := []string{"Venere ", " Terra", " Marte"} //pianeti con degli spazi
	hyperspace(planets)
	fmt.Println(strings.Join((planets),""))

}
