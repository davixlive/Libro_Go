package main

import "fmt"

func main() {
	fmt.Println("C'è un cartello vicino all'ingresso, che dice 'Vietato ai minori'")
	var age = 41
	var minor = age < 18
	fmt.Printf("Ho %v anni, sono minorenne? %v\\n", age, minor )
}
