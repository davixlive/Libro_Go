package main

import "fmt"

//mostra la lunghezza, la capacità e il contenuto della slice
func dump (label string, slice [] string){
	fmt.Printf("%v: lungh. %v, capacità %v %v\n",label, len(slice), cap(slice),slice)
}
func main() {
	dwarfs := []string {"Cenere","Plutone","Haumea","Makemake","Eris"}
	dump("nani",dwarfs)
	dump("nani[1:2]",dwarfs[1:2])

}
