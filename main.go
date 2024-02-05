package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Oficial do São Paulo", Preco: 349.99, Quantidade: 15},
		{"Tênis", "Corrida Meia Maratona", 799.99, 22},
		{"Bicicleta", "Bike da praia", 1289.56, 12},
		{"Meia", "Meia com pressão", 129.89, 5},
	}

	temp.ExecuteTemplate(w, "index", produtos)
}
