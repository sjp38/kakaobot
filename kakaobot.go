package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "received %s %s",
		r.Method, html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/kakaobot/", handleHTTP)

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Printf("buy\n")
}
