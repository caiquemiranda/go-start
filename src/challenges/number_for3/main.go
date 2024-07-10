package main

import "fmt"

func main() {
	fmt.Println("Numeros divisiveis por 3: ")
	for i := 1; i <= 100; i++ {
		
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
