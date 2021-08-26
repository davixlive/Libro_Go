package main

import (
	"fmt"

)

func main() {
	launch := false
	launchText := fmt.Sprintf("%v",launch)
	fmt.Println("Pronti per il lancio:",launchText)
	var yesNo string
	if launch{
		yesNo = "sì"

	}else {
		 yesNo = "No"
	}
	fmt.Println("Pronti per il lancio", yesNo)
}
