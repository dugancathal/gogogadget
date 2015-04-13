package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port          = ""
)

func init() {
	port = os.Getenv("PORT")
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`pong`))
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := []string{"TJ", "Greg"}
		body, _ := json.Marshal(map[string][]string{"users": users})
		w.Write(body)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
