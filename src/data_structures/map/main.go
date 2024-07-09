package main

import "fmt"

func main() {
	/*
		x := make(map[string]int)
		x["chave"] = 10

		fmt.Println(x["chave"])
	*/

	/*
		y := make(map[int]int)
		y[1] = 10
		y[2] = 20
		y[3] = 30

		fmt.Println(y[1], y[2], y[3])
	*/

	element := make(map[string]string)
	element["O"] = "Oxigênio"
	element["H"] = "Hidrogênio"
	element["He"] = "Hélio"
	element["Li"] = "Lítio"
	element["Be"] = "Berílio"
	element["B"] = "Boro"
	element["C"] = "Carbono"
	element["N"] = "Nitrogenio"

	fmt.Println(element["H"], element["Li"])

}
