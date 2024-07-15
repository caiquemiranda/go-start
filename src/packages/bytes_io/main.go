package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func tipe1() {
	if _, err := io.WriteString(os.Stdout, "Helleo World"); err != nil {
		log.Fatal(err)
	}
}

func tipe2() {
	msg := []byte("Rei Leão World!")
	err := ioutil.WriteFile("hello", msg, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func tipe3() {
	fmt.Printf("%s", bytes.Title([]byte("he royal highess")))
}
func main() {
	tipe1()
	tipe2()
	tipe3()
}
