package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func request(w http.ResponseWriter, r *http.Request) {
	dest := os.Getenv("DEST")
	message, err := http.Get(dest)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}
	w.Write([]byte("segundo\n"))
	defer message.Body.Close()
	body, err := ioutil.ReadAll(message.Body)
	if err != nil {
		print(err)
		return
	}
	w.Write([]byte(body))
}

func start() {
	http.HandleFunc("/second", request)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
