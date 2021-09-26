package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("post")))
	r.ParseForm()
	fmt.Println(r.Header["Content-Type"])
	str, err := json.Marshal(r.Header["Content-Type"])
	if err != nil {
		log.Fatal(err)
	}
	w.Write(str)
}

func main() {
	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	server.Handler = mux
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
