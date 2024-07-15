package main

import "fmt"

func main() {
	x := 5
	inicial(&x)
	fmt.Println(x)
}

func inicial (xPtr *int){
	*xPtr = 0
}