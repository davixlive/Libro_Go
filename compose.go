package main

import "fmt"

type report struct {
	sol int
	temperature temperature
	location location
}
type temperature struct {
	high, low celsius
}
type location struct {
	lat,long float64
}

type celsius float64

func main() {
	bradburry := location{-4.5895, 137.4417}
	t := temperature {high: 1.0, low: -78.0}
	report := report{sol : 15, temperature: t , location: bradburry}
	fmt.Printf("%+v\n",report)
	fmt.Printf("Un tiepido %v° C\n",report.temperature.high)

}
