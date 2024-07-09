package main

import "fmt"

func main() {
	numbersOne := make([]float64, 10)
	fmt.Println(numbersOne)

	arr := [7]float64{1, 3, 5, 7, 9, 11, 12}
	slice := arr[3:5]
	fmt.Println(slice)

	arrTwo := []int{1, 2, 3}
	arrThree := append(arrTwo, 4, 5)
	fmt.Println(arrTwo)
	fmt.Println(arrThree)

	arrFour := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arrFive := make([]int, 3)
	copy(arrFive, arrFour)
	fmt.Println(arrFive)

}
