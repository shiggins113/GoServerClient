package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(create string) *sql.DB {
	db1, err := sql.Open("mysql", create)
	if err != nil {
		fmt.Println("db.go error")
		log.Fatal(err)
	}

    file, err := ioutil.ReadFile("createDB.sql")
    if err != nil {
		log.Fatal(err)
	}
	sqlQuery := strings.TrimSpace(string(file))

    requests := strings.Split(sqlQuery, ";")

    for _, sqlStatement := range requests {
		if len(strings.TrimSpace(sqlStatement)) == 0 {
			continue
		}   

        _, err := db1.Exec(sqlStatement)
		if err != nil {
			fmt.Println("ERROR")
			log.Fatal(err)
		}

	}

	return db1
}