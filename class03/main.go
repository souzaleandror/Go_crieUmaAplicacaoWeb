package main

import ( 
	"net/http"
	"routes"
)

func main() {
	// db := connectComBancoDeDados()
	// defer db.Close()
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

