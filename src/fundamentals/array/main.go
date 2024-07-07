package main

import "fmt"

func main() {
	var x [5]int

	x[0] = 1
	x[1] = 23
	x[2] = 45

	fmt.Println(x)

	var y [10]int

	y[3] = 10
	y[4] = 20
	y[6] = 90
	y[7] = 30

	fmt.Println(y)
}
