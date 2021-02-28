package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"

	"net/http"
)

type Data struct {
	Amount int
}

var mux = http.NewServeMux()


func getall(w http.ResponseWriter, r *http.Request){


	version := 1
	var e Data
	err1 := json.NewDecoder(r.Body).Decode(&e)
	if err1 != nil {
		panic(err1)
	}

	connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("INSERT INTO reqests (data, version) values ($1,$2)", e.Amount, version)

}





func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux


	mux.HandleFunc("/getall", getall)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	//dsadsad
}