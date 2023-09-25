package main

import ( 
	"database/sql"
	"net/http"
	"html/template"
	_ "github.com/lib/pq"
)

func connectComBancoDeDados() *sql.DB {
	//conexao := "user dbname password host sslmode"
	conexao := "user=postgres dbname=alura_loja password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil{
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Id int
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// db := connectComBancoDeDados()
	// defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	// produtos := []Produto{
	// 	{Nome:"Camiseta", Descricao:"Azul, bem bonito", Preco:30, Quantidade: 5},
	// 	{"Tenis", "Confortavel", 89, 3},
	// 	{"Fone", "Muito bom", 189, 4},
	// }

	db := connectComBancoDeDados()

	selectTodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosOsProdutos.Next() {
		var id, quantidade int 
		var nome, descricao string
		var preco float64

		err = selectTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		p.Id = id

		produtos = append(produtos, p)
	}

	defer db.Close()

	temp.ExecuteTemplate(w, "Index", produtos);
}