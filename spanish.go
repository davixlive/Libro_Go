package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"
	fmt.Println(len(question),"byte")
	fmt.Println(utf8.RuneCountInString(question),"rune")
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("Prima runa: %c - %v byte", c, size )

}
