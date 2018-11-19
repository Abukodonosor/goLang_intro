package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //extract package and allowe driver comunication with databse/sql package
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root12"
	password = "root12"
	dbname   = "zoo"
)

type Toy struct {
	id       int
	name     string
	lastname string
	credits  int
}

//https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.4.html
//	This is link where you can check sql syntax
//

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var lastInsertId int
	var name string
	var credits int
	// var rows sql.Rows
	var rows *sql.Rows

	err = db.QueryRow("INSERT INTO public.test(name, lastname, credits) VALUES($1,$2,$3) returning id,name,credits;", "Jehovini", "Svedoci", 12).Scan(&lastInsertId, &name, &credits)
	checkErr(err)
	fmt.Printf("last inserted id :%d name: %s credits:%d \n", lastInsertId, name, credits)

	var sqlQuery string
	sqlQuery = "SELECT * FROM public.test"

	rows, err = db.Query(sqlQuery)
	checkErr(err)

	// fmt.Printf("%s", *rows[0])

	for rows.Next() {

		// var objectNew *Toy
		var xid int
		var xuser string
		var xcredits int
		var xlastname string
		err = rows.Scan(&xid, &xuser, &xlastname, &xcredits)

		newwqw := Toy{xid, xuser, xlastname, xcredits}
		checkErr(err)
		fmt.Printf("%s", rows)
		fmt.Println("id | username | lastname | credits ")
		fmt.Printf("%d | %s | %s | %d\n", xid, xuser, xlastname, xcredits)
	}

	fmt.Println("Successfully connected!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
