package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	urlconn := "ismael:adm123@/golang?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", urlconn)

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Opened Conn")

	lines, erro := db.Query("select * from users")

	if erro != nil {
		log.Fatal(erro)
	}
	defer lines.Close()

	fmt.Println(lines)

}
