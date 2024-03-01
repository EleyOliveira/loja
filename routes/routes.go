package routes

import (
	"loja/controllers"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/inclusao", controllers.Inclusao)
	http.HandleFunc("/insert", controllers.Incluir)
}
