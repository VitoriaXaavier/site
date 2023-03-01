package routes

import (
	"net/http"
	"github.com/VitoriaXaavier/site/controle"
)

func CarregaRotas () {
	http.HandleFunc("/", controle.Index)
	http.HandleFunc("/new", controle.New)
	http.HandleFunc("/insert", controle.Insert)
	http.HandleFunc("/delete",controle.Deleta)
	http.HandleFunc("/edit", controle.Edita)
	http.HandleFunc("/update", controle.Update)
}


