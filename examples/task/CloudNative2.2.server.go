package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
		for _, str := range v[1:] {
			w.Header().Add(k, str)
		}
	}
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	log.Println(r.Host, 200)
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	log.SetOutput(os.Stdout)
	err := os.Setenv("VERSION", "1")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/healthz", HealthzHandler)
	server.Handler = mux
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
