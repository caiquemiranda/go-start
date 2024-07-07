package main

import "fmt"

func main() {
	x := 3

	if x == 2 || x == 3 {
		fmt.Println("X é 2 ou 3")
	} else {
		fmt.Println("X não 2 nem 3")
	}

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("X é divisível por 2 e 3")
	} else {
		fmt.Println("X não é divisível por 2 e nem 3")
	}
}
