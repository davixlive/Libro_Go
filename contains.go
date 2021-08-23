package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Ti trovi in una caverna buia.")
	var command = "vai fuori"
	var exit = strings.Contains(command, "fuori")
	fmt.Println("Esci dalla caverna:", exit) // Scrive Uscite dalla caverna: true
}
