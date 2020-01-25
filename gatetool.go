package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func baseAccess(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome lmao !")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", baseAccess)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
