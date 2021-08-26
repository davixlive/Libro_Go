package main

import "fmt"

func main() {
	var command = "vai est"
	if command == "vai est"{
		fmt.Println("Ti dirigi verso la montagna")
	}else if command == "vai dentro"{
		fmt.Println("Entri nella caverna, nella quale trascorrerai il resto della tua vita")
	}else{
		fmt.Println("Non ho capito")
	}
}
