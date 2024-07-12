package main

import "fmt"

func media(lista []float64) float64{
	total := 0.0

	for _, valor := range lista {
		total += valor
	}

	return total / float64(len(lista))
}

func main() {
	lista := []float64{4.5, 5.6, 7.2, 4.1, 9, 1}
	media := media(lista)
	fmt.Printf("A média das notas: %.2f\n", media)

}

/*
total := 0.0

	for _, valor := range lista {
		total += valor
	}
	media := total / float64(len(lista))
	fmt.Println("A média das notas: ", media)
*/
