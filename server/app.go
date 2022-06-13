package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// applicationが利用するrouterとdbの参照をexposeする
type App struct {
	Router *mux.Router
	DB *sql.DB
}


// DBへの接続のための初期化メソッド
func (app *App) Initialize(user, password, dbname string) {
	connectionString := 
		fmt.Sprintf("user=%s, password=%s, dbname=%s sslmode=disable", user, password, dbname)
	var err error

	app.DB, err = sql.Open("postgres",connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
}


// アプリケーションをスタートする
func (app *App) Run(addr string){}

