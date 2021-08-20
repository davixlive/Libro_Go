package main

import "fmt"
/*
	Il sistema di trasporto interplanetario SpaceX non è dotato della velocità warp della
	nave spaziale Enterprise, ma è in grado di portavi su Marte alla rispettabile velocità di 100800 km/h.
	Un'ambiziosa data di lancuo nel gennaio 2025 troverebbe Marte e la Terra a una distanza di 96300000 km.
	Quanti giorni sarebbero necessari per raggiungere Marte?
 */
func main() {
	const velocitaSpaceX = 100800 // Km/h
	var distanza = 96300000
	fmt.Println("Per raggiungere Marte si impiegherebbero ", distanza/velocitaSpaceX, "giorni.")

}
