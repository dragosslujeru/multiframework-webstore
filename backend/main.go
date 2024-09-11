package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting WebStore backend")

	fs := http.FileServer(http.Dir("../vanilla-js"))
	mux := http.NewServeMux();
	mux.Handle("/vanilla-js/", http.StripPrefix("/vanilla-js", fs))

	e := http.ListenAndServe(":8080", mux)
	if (e != nil) {
		log.Fatalf("Error while starting WebStore backend %e", e)
	}
	
}