package main

import (
	"loja/routes"
	"net/http"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
