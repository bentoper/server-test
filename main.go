package main

import (
	"fmt"
	"net/http"
)

func answerRequest(w http.ResponseWriter, r *http.Request) {
	message := "Uma resposta\n"
	fmt.Fprint(w, "Oi\n")
	w.Write([]byte(message))
}

func main() {
	go start()
	http.HandleFunc("/first", answerRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
