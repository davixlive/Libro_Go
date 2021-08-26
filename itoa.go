package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown := 10
	str := "Lancio fra " + strconv.Itoa(countdown) + " secondi"
	fmt.Println(str)
}
