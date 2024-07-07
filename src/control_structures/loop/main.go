package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Print(horas, " Horas ")

		for minutos := 0; minutos <= 59; minutos++ {
			fmt.Print(minutos, " minutos ")
			fmt.Println()

		}

	}

}
