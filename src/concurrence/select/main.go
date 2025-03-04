package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() { //Recebemos os valores "um"e depois "dois"conforme o esperado. 2 segundos,
		//pois o 1 e o 2 segundos são Sleeps executados simultaneamente.
		c1 <- "um"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "dois"
	}()

	for i := 0; i < 2; i++ {
		select { //Usaremos select para aguardar esses
		//dois valores simultaneamente, imprimindo cada um à medida que chegam.
		case msg1 := <-c1:
			fmt.Println("receba", msg1)
		case msg2 := <-c2:
			fmt.Println("receba", msg2)
		}
	}
}
