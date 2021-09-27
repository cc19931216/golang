package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wc sync.WaitGroup

func Client(client *http.Client, url string) {
	defer wc.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("1", "a")
	req.Header.Add("1", "b")
	req.Header.Set("2", "c")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", res.Header)
	log.Printf("%v", res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data)
}

func main() {
	client := &http.Client{}
	rootURL := "http://127.0.0.1:8080"
	wc.Add(1)
	go Client(client, rootURL)
	wc.Wait()

	rootURL = "http://127.0.0.1:8080/healthz"
	wc.Add(1)
	go Client(client, rootURL)
	wc.Wait()
}
