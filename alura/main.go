package main

import ( 
	"net/http"
	"html/template"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{Nome:"Camiseta", Descricao:"Azul, bem bonito", Preco:30, Quantidade: 5},
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Muito bom", 189, 4},
	}

	temp.ExecuteTemplate(w, "Index", produtos);
}