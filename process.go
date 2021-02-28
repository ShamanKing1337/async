package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Request struct{
	Id int
	Amount int
	Version int

	//Isproceeded bool
}


func report(){
	for true{
		connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		tr, _ := db.Begin()
		rows := tr.QueryRow("SELECT id, data, version FROM  reqests where isproceeded is NULL LIMIT 1 for update SKIP LOCKED")
		if err != nil{
			fmt.Println(err)
		}
		p := Request{}
		rows.Scan(&p.Id, &p.Amount, &p.Version)
		time.Sleep(10 * time.Second)


		_, err =tr.Exec("UPDATE reqests SET isproceeded = 'True' where id = $1", p.Id)
		if err != nil{

			fmt.Println(err)
		}
			fmt.Println(p.Id)


		tr.Commit()



	}

}

func main(){

	report()

	//time.Sleep(100 * time.Second)
}

