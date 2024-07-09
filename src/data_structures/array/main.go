package main

import "fmt"

func main() {
	var numbers [5]float64

	numbers[0] = 3.1
	numbers[1] = 7.2
	numbers[2] = 6.5
	numbers[3] = 8.5
	numbers[4] = 10.0

	fmt.Println(numbers)

	var total float64 = 0

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]

	}
	fmt.Println("A média dos valores do array é: ", total/float64(len(numbers)))

}
