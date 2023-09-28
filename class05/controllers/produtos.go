package controllers

import (
	"html/template"
	"log"
	"models"
	"net/http"
	"strconv"
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

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditarProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		idConvertido, err3 := strconv.Atoi(id)
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

		if err3 != nil {
			log.Println("Error ao converter Id: ", err3)
		}

		models.AtualizarNovoProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConvertido)
	}

	http.Redirect(w, r, "/", 301)
}
