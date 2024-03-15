package routes

import (
	"loja/controllers"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/inclusao", controllers.Inclusao)
	http.HandleFunc("/insert", controllers.Incluir)
	http.HandleFunc("/delete", controllers.Deletar)
	http.HandleFunc("/edicao", controllers.Edicao)
	http.HandleFunc("/edit", controllers.Editar)
}
