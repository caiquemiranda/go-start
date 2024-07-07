package main

import "fmt"

func main() {
	var name string = "Caique"
	var age int = 31
	var city string = "Salvador"
	var version float32 = 0.1

	fmt.Printf("Meu nome é : %s e minha idade é : %d", name, age)
	fmt.Printf("\nMinha cidade é : %s e a minha versão do software é : %.2f\n", city, version)
}