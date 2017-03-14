package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

type keyboard struct {
	Type string `json:"type"`
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, "received %s %s",
		r.Method, html.EscapeString(r.URL.Path))
	if r.Method == "GET" && r.URL.Path == "/kakaobot/keyboard" {
		resp, err := json.Marshal(keyboard{
			Type: "keyboard"})
		if err != nil {
			log.Fatal("Failed to marshal keybaord: %s", err)
		}
		fmt.Fprintf(w, string(resp))
		return
	}
}

func main() {
	http.HandleFunc("/kakaobot/", handleHTTP)

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Printf("buy\n")
}
