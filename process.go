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


func main(){

	for true{
		connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		rows,err := db.Query("SELECT id, data, version FROM  reqests where isproceeded is NULL LIMIT 1 FOR UPDATE")
		requests := []Request{}

		for rows.Next(){
			p := Request{}
			err := rows.Scan(&p.Id, &p.Amount, &p.Version)
			if err != nil{
				fmt.Println(err)
				continue
			}
			requests = append(requests, p)
			db.Exec("UPDATE reqests SET isproceeded = 'True' where id = $1", p.Id)
			fmt.Println(p.Id)

		}

		fmt.Println(requests)

		time.Sleep(10 * time.Second)
	}




}

