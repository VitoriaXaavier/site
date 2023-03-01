package main

import (
	"net/http"

	"github.com/VitoriaXaavier/site/routes"
)

func main() {

	routes.CarregaRotas()
	
	http.ListenAndServe(":8000", nil)

}


