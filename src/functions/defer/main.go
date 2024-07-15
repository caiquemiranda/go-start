package main

import "fmt"

func dia1() {
	fmt.Println("Domingo")
}

func dia2() {
	fmt.Println("Segunda-Feira")
}

func dia3() {
	fmt.Println("Terça-Feira")
}

func main() {

	dia2()
	defer dia1()
	dia3()
}
