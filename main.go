package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	Yellow = "\033[1;33m%s\033[0m\n"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	data := "Incredible India!"
	//dataBytes := []byte(data)

	if r.Method == "GET" {
		n, err := fmt.Fprint(w, data)
		if err != nil && n <= 0 {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Something went wrong")
		}
	}
}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Printf(Yellow, "Listening on port 20001")
	err := http.ListenAndServe("localhost:20001", nil)
	if err != nil {
		log.Fatal(err)
	}

}
