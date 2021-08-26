package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
func mesaureTemperature (samples int, sensor func() kelvin){
	for i := 0; i < samples ; i++{
		k := sensor()
		fmt.Printf("%v° K\n",k)
		time.Sleep(time.Second)
	}
}
func fakeSensor() kelvin{
	return kelvin(rand.Intn(151) + 150)
}
func main() {
	mesaureTemperature(3,fakeSensor)
}
