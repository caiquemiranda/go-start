package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		
		if i == 4 {
			fmt.Println(i, "LOOP")
		} else if i == 7 {
			fmt.Println(i, "é igual a 7")
			continue
		} else if i == 9 {
			fmt.Println(i, "é igual a 9")
			break
		}
		fmt.Println(i)
	}
}
