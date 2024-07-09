package main

import "fmt"

func main() {
	var numbers [5]float64

	numbers[0] = 3.1
	numbers[1] = 7.2
	numbers[2] = 6.5
	numbers[3] = 8.5
	numbers[4] = 5.4

	fmt.Print(numbers)

	var total float64 = 0.0

	for i := 0; 0 < 5; i++ {
		total += numbers[i]
		fmt.Println(total)
	}


}
