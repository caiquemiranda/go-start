package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf(" O %q artigo foi escrtito por %s", a.Title, a.Author)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Primeiros passos em Go",
		Author: "Caique Miranda",
	}
	Print(a)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
