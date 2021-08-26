package main

import (
	"fmt"
	"math/big"
)

/*
	Giorni necessari per raggiungere la galassia di Andromeda
	package BIG
 */
func main() {
	lightspeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	distance := new(big.Int)
		distance.SetString("24000000000000000000",10)
	fmt.Println("La galssia di Andromeda è a", distance, "km.")
	seconds := new(big.Int)
	seconds.Div(distance,lightspeed)
	days := new(big.Int)
	days.Div(seconds,secondsPerDay)
	fmt.Println("Ci vogliono",days,"giorni, alla velocità della luce.")

}
