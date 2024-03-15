package controllers

import (
	"html/template"
	"log"
	"loja/models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.ConsultaProdutos()
	temp.ExecuteTemplate(w, "index", produtos)
}

func Inclusao(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "inclusao", nil)
}

func Incluir(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
	if err != nil {
		log.Println("Formato inválido de preço", err)
	}

	quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
	if err != nil {
		log.Println("Formato inválido de quantidade", err)
	}

	models.Incluir(nome, descricao, preco, quantidade)

	http.Redirect(w, r, "/", 301)

}

func Edicao(w http.ResponseWriter, r *http.Request) {
	produto := models.ConsultaProduto(r.URL.Query().Get("id"))
	temp.ExecuteTemplate(w, "edicao", produto)
}

func Editar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Formato inválido de preço", err)
		}
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Formato inválido de quantidade", err)
		}

		models.Editar(id, nome, descricao, preco, quantidade)
		http.Redirect(w, r, "/", 301)
	}

}

func Deletar(w http.ResponseWriter, r *http.Request) {
	models.Deletar(r.URL.Query().Get("id"))
	http.Redirect(w, r, "/", 301)
}
