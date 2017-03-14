package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type keyboard struct {
	Type string `json:"type"`
}

type message struct {
	Userkey string `json:"user_key"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type resptext struct {
	Text string `json:"text"`
}

type response struct {
	Message resptext `json:"message"`
}

type user struct {
	UserKey string `json:"user_key"`
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("received %s %s\n",
		r.Method, html.EscapeString(r.URL.Path))
	if r.Method == "GET" && r.URL.Path == "/kakaobot/keyboard" {
		resp, err := json.Marshal(keyboard{
			Type: "text"})
		if err != nil {
			log.Fatal("Failed to marshal keybaord: %s", err)
		}
		fmt.Fprintf(w, string(resp))
		return
	}

	if r.Method == "POST" && r.URL.Path == "/kakaobot/message" {
		body, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			log.Fatal("Failed to read body of /message: %s", err)
		}
		log.Printf("body: %s\n", string(body))
		var msg message
		if err := json.Unmarshal(body, &msg); err != nil {
			log.Fatal("Failed to unmarshal body of /message: %s %s", err, string(body))
		}
		resp, err := json.Marshal(response{
			Message: resptext{
				Text: msg.Content}})
		if err != nil {
			log.Fatal("Failed to marshal response: %s", err)
		}
		log.Printf("send %s\n", string(resp))
		fmt.Fprintf(w, string(resp))
		return
	}

	if r.Method == "POST" && r.URL.Path == "/kakaobot/friend" {
		body, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			log.Fatal("Failed to read body of /friend: %s", err)
		}
		var usr user
		if err := json.Unmarshal(body, &usr); err != nil {
			log.Fatal("Failed to unmarshal body of /friend: %s %s", err, string(body))
		}
		fmt.Fprintf(os.Stdout, "Fried %s joined!\n", usr.UserKey)
		return
	}

	if r.Method == "DELETE" {
		tokens := strings.Split(r.URL.Path, ":")
		if tokens[0] == "/kakaobot/friend/" {
			fmt.Fprintf(os.Stdout, "Friend %s banned me!\n", tokens[1])
		}
		if tokens[0] == "/kakaobot/chat_room/" {
			fmt.Fprintf(os.Stdout, "Friend %s leaved chat room!\n", tokens[1])
		}
		return
	}
}

func main() {
	http.HandleFunc("/kakaobot/", handleHTTP)

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Printf("buy\n")
}
