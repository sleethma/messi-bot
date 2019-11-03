package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// MainIndex returns general message
func MainIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from mlabouardy :)")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", TopIndex)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}