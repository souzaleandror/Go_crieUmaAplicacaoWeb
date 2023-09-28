package db

import ( 
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectComBancoDeDados() *sql.DB {
	//conexao := "user dbname password host sslmode"
	conexao := "user=postgres dbname=alura_loja password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil{
		panic(err.Error())
	}

	return db
}