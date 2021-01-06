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
	/*client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}*/
	message, err := http.Get(dest)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
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
