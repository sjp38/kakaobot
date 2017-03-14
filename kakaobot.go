package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, "received %s %s",
		r.Method, html.EscapeString(r.URL.Path))
	if r.URL.Path == "/kakaobot/keyboard" {
		fmt.Fprintf(w, "{\n\"type\":\"text\"\n}")
		return
	}
}

func main() {
	http.HandleFunc("/kakaobot/", handleHTTP)

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Printf("buy\n")
}
