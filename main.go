package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

type url struct {
	Orig_URL    string `json:"orig_url"`
	Reduced_URL string `json:"reduced_url"`
	Date        string `json:"date"`
}

type urls []url

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe("2525", router))
}
