package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown, err := strconv.Atoi("10")
	if err != nil { //Il valore nil per err significa che non si è verificato alcun errore
		//Qualcosa è andato storto
	}
	fmt.Println(countdown)
}
