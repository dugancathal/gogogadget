package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	port  = os.Getenv("PORT")
	dbUrl = os.Getenv("DATABASE_URL")
)

func init() {
	time.Sleep(3 * time.Second)
}

func main() {
	uri, _ := url.Parse(dbUrl)
	databaseUrl := fmt.Sprintf("%s@tcp(%s)%s", uri.User, uri.Host, uri.Path)
	fmt.Println("Database connection to:", databaseUrl)
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		fmt.Println("*******Database connection error*******")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("*******Database connection error*******")
		panic(err)
	}

	fmt.Println("*******Executing startup SQL*******")
	db.Exec("CREATE TABLE IF NOT EXISTS users ( id INTEGER NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL, PRIMARY KEY (id) ) ENGINE=InnoDB, CHARSET=utf8;")
	db.Exec("TRUNCATE TABLE users;")
	db.Exec(`INSERT INTO users (name) VALUES ("TJ"), ("Greg");`)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`pong`))
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name FROM users;")
		if err != nil {
			log.Fatal(err)
		}
		users := map[string]int64{}

		defer rows.Close()
		for rows.Next() {
			var id int64
			var name string
			err = rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			users[name] = id
		}

		body, err := json.Marshal(users)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(body)
	})

	fmt.Println("*******Server starting*******")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
