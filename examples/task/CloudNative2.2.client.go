package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var wc sync.WaitGroup

func HelloClient(client *http.Client, url string, method string) {
	defer wc.Done()
	req := &http.Request{}
	var err error
	if method == "POST" {
		reqData := "name=ali&age=19"
		req, err = http.NewRequest(method, url, strings.NewReader(reqData))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		req, err = http.NewRequest(method, url, nil) //GET
		if err != nil {
			log.Fatal(err)
		}
	}
	//设置Content-Type很重要
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rep, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(rep.Body)
	rep.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data)
}

func main() {
	client := &http.Client{}
	rootURL := "http://127.0.0.1:8080"
	wc.Add(1)
	go HelloClient(client, rootURL, "POST")
	wc.Wait()
}
