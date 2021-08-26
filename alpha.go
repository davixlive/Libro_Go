package main

import "fmt"

/*
	Giorni necessari per raggiungere Alpha Centauri
 */
func main() {
	const lightspeed = 299792 //km/s
	const secondsPerDay = 86400
	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri è a", distance, "km.")
	days := distance / lightspeed / secondsPerDay
	fmt.Println("Ci vogliono", days, "giorni, alla velcoità della luce.")
}
