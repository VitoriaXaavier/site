package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descrição  string
	Preço      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto {

		{Nome: "Shorts", Descrição: "Preta-Lisa", Preço: 29.9, Quantidade: 5},
		{"Tenis", "Branco", 129.9,3},
		{"Fone","Bluetooth", 155,2},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
