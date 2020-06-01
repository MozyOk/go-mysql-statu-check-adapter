package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	user     = flag.String("user", "", "database user name")
	password = flag.String("password", "", "database password")
	database = flag.String("database", "", "database")
	query    = flag.String("query", "", "test query")
	address  = flag.String("address", "localhost:8080", "address listen on")
)

func main() {
	flag.Parse()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", *user, *password, *database))
	if err != nil {
		fmt.Printf("Error opening Database: %v", err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		_, err := db.Exec(*query)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("OK"))
		return
	})

	http.ListenAndServe(*address, nil)
}
