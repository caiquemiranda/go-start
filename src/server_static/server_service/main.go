package main

import (
    "fmt"
	"net/http"
)

func ola(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Olá")
}

func cabecalhos(w http.ResponseWriter, r *http.Request){
	for nome , cabecalhos := range r.Header{
		for _, c := range cabecalhos{
			fmt.Fprintf(w, "%s: %s\n", nome, c)
		}
	}
}


func main(){
    http.HandleFunc("/ola", ola)
    http.HandleFunc("/cabecalhos", cabecalhos)
    fmt.Println("Servidor rodando na porta 8080...")
    http.ListenAndServe(":8080", nil)
}