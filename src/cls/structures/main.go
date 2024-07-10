package main

import "fmt"

type Pessoa struct {
	nome string	
    idade int
}


// coleções de campos
func main() {
	// type Pessoa struct {}
	fmt.Println(Pessoa{"Caique", 31})
	fmt.Println(Pessoa{"Maria", 28})
	fmt.Println(Pessoa{"João", 35})
}
