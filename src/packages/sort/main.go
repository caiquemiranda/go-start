package main

import (
    "fmt"
    "sort"
)

type Dados struct {
    nome string
    idade int
}

type ParaNome []Dados

func (ps ParaNome) Len() int { 
    return len(ps)
}

func (ps ParaNome) Less(i, j int) bool {
    return ps[i].nome < ps[j].nome
}

func (ps ParaNome) Swap(i, j int) {
    ps[i], ps[j] = ps[j], ps[i]
}

func main() {
    kids := []Dados{
        {"Alice", 8},
        {"Bob", 5},
        {"Charlie", 12},
    }
    sort.Sort(ParaNome(kids))
    fmt.Println(kids)
}