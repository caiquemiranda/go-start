package main

import "fmt"

func main() {
	x:=0

	number := func ()int  {
		x++
        return x
		
	}
	fmt.Println(number()) // Output: 1
	fmt.Println(number()) // Output: 2
	fmt.Println(number()) // Output: 3
	fmt.Println(number()) // Output: 4
	fmt.Println(number()) // Output: 5

}