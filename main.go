package main

import (
	"fmt"
	"net/http"
	"os"
)

func answerRequest(w http.ResponseWriter, r *http.Request) {
	dest := "Dest: " + string(os.Getenv("DEST")) + "\n"
	message := "Uma resposta\n"
	fmt.Fprint(w, dest)
	w.Write([]byte(message))
}

func main() {
	go start()
	http.HandleFunc("/first", answerRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
