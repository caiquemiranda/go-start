package main

import (
    "fmt"
    "strings"
)

func main() {
	fmt.Println(strings.Contains("Computer", "ter"))
	fmt.Println(strings.Contains("Terra", "mor"))
	fmt.Println(strings.Contains("Computer", "Ter"))
}