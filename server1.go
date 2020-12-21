package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func request(w http.ResponseWriter, r *http.Request) {
	message, err := http.Get("http://localhost:8080/first")
	if err != nil {
		fmt.Println(err)
		return
	}

	//defer message.Body.Close()
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
