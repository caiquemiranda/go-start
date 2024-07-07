package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var value = 1

	if value > 0{
		fmt.Println("O valor é maior que zero")
	}

	switch value {
		case 1:
            fmt.Println("O valor é 1")
        case 2:
            fmt.Println("O valor é 2")
        default:
            fmt.Println("O valor não é 1 nem 2")
		}
}