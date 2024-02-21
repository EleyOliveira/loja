package routes

import (
	"loja/controllers"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controllers.Index)
}
