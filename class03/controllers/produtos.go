package controllers

import(
	"net/http"
	"html/template"
	"models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		precoConvertido, err := strconv.ParseFloat(preco, 64)
		quantidadeConvertido, err2 := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Error ao converter Preco: ", err)
		}

		if err2 != nil {
			log.Println("Error ao converter Quantidade: ", err2)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)
	}

	http.Redirect(w, e, "/", 301)
}