package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)

}
func oops() error {
	return MyError{
		time.Date(19989, 3, 15, 22, 30, 0, 0, time.UTC),
		"O arquivo do sistema desapareceu!",
	}
}
func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
