package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
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

func msgFor(tokens []string) string {
	if tokens[0] == "ex" {
		if len(tokens) < 2 {
			return "You forgot command."
		}
		allowed, ok := executables[tokens[1]]
		if !ok || !allowed {
			return "It cannot be executed."
		}

		out, err := exec.Command("./"+tokens[1], tokens[2:]...).Output()
		if err != nil {
			return "Failed to execute the command."
		}
		return string(out)
	}
	// Just echo received message.
	return strings.Join(tokens, "... ") + "...???"
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("received %s %s\n",
		r.Method, html.EscapeString(r.URL.Path))
	if r.Method == "GET" && r.URL.Path == "/kakaobot/keyboard" {
		resp, err := json.Marshal(keyboard{
			Type: "text"})
		if err != nil {
			log.Fatal("Failed to marshal /keybaord: %s", err)
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
			log.Fatal("Failed to unmarshal body of /message: %s",
				err)
		}
		message := rawMsgToMessageKey(msg.Content)
		if message == "exception" {
			message = msg.Content
		}
		rmsg := msgFor(strings.Fields(message))
		resp, err := json.Marshal(response{
			Message: resptext{
				Text: rmsg}})
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
		log.Printf("body: %s\n", string(body))
		var usr user
		if err := json.Unmarshal(body, &usr); err != nil {
			log.Fatal("Failed to unmarshal body of /friend: %s",
				err)
		}
		log.Printf("Friend %s joined!\n", usr.UserKey)
		return
	}

	if r.Method == "DELETE" {
		tokens := strings.Split(r.URL.Path, ":")
		if tokens[0] == "/kakaobot/friend/" {
			log.Printf("Friend %s banned me!\n", tokens[1])
		}
		if tokens[0] == "/kakaobot/chat_room/" {
			log.Printf("Friend %s leaved chat room!\n", tokens[1])
		}
		return
	}
}

var rawMsgToMsgKeyMap = map[string]string{
	"hi": "hi",
	"hello": "hi",
}

func rawMsgToMessageKey(rawMessage string) string {
	key, ok := rawMsgToMsgKeyMap[rawMessage]
	if !ok {
		return "exception"
	}
	return key
}

func loadExecutables(filepath string) bool {
	c, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("failed to read executables from file: %s\n", err)
		return false
	}
	if err := json.Unmarshal(c, &executables); err == nil {
		fmt.Printf("failed to unmarshal executables: %s\n", err)
		return false
	}

	return true
}

func loadMsgToKey(filepath string) bool {
	c, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("failed to read messages to keys: %s\n", err)
		return false
	}
	if err := json.Unmarshal(c, &rawMsgToMsgKeyMap); err != nil {
		fmt.Printf("failed to unmarshal messages: %s\n", err)
		return false
	}
	return true
}

var executables = map[string]bool{
	"ls": false,
}

func saveMsgToKey(filepath string) {
	bytes, err := json.Marshal(rawMsgToMsgKeyMap)
	if err != nil {
		fmt.Printf("failed to marshal messages: %s\n", err)
		return
	}

	if err := ioutil.WriteFile(filepath, bytes, 0600); err != nil {
		fmt.Printf("failed to write messages: %s\n", err)
		return
	}
}

func saveExecutables(filepath string) {
	bytes, err := json.Marshal(executables)
	if err != nil {
		fmt.Printf("failed to marshal executables: %s\n", err)
		return
	}

	if err := ioutil.WriteFile(filepath, bytes, 0600); err != nil {
		fmt.Printf("failed to writre executables: %s\n", err)
		return
	}
}

func main() {
	msgtoKeyFile := "msg_to_key.json"
	if !loadMsgToKey(msgtoKeyFile) {
		saveMsgToKey(msgtoKeyFile)
	}

	exeFile := "executables.json"
	if !loadExecutables(exeFile) {
		saveExecutables(exeFile)
	}

	http.HandleFunc("/kakaobot/", handleHTTP)

	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Printf("buy\n")
}
