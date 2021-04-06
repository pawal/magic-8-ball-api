package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Magic struct {
	Id     int    `json:"Id"`
	Answer string `json:"Answer"`
	Type   string `json:"Type"`
}

var MagicAnswers []Magic

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setJson(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/ reached")
	fmt.Fprintf(w, "Magic 8-ball API!")
}

func returnRandomAnswer(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(MagicAnswers))
	fmt.Printf("/magic answered with %d: %s\n", randomIndex, MagicAnswers[randomIndex].Answer)
	setJson(&w)
	enableCors(&w)
	json.NewEncoder(w).Encode(MagicAnswers[randomIndex])
}

func returnAnswers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/completemagic answered")
	setJson(&w)
	enableCors(&w)
	json.NewEncoder(w).Encode(MagicAnswers)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/magic", returnRandomAnswer)
	http.HandleFunc("/completemagic", returnAnswers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	MagicAnswers = []Magic{
		{0, "It is certain.", "affirmative"},
		{1, "It is decidedly so.", "affirmative"},
		{2, "Without a doubt.", "affirmative"},
		{3, "Yes – definitely.", "affirmative"},
		{4, "You may rely on it.", "affirmative"},
		{5, "As I see it, yes.", "affirmative"},
		{6, "Most likely.", "affirmative"},
		{7, "Outlook good.", "affirmative"},
		{8, "Yes.", "affirmative"},
		{9, "Signs point to yes.", "affirmative"},
		{10, "Reply hazy, try again.", "neutral"},
		{11, "Ask again later.", "neutral"},
		{12, "Better not tell you now.", "neutral"},
		{13, "Cannot predict now.", "neutral"},
		{14, "Concentrate and ask again.", "neutral"},
		{15, "Don't count on it.", "negative"},
		{16, "My reply is no.", "negative"},
		{17, "My sources say no.", "negative"},
		{18, "Outlook not so good.", "negative"},
		{19, "Very doubtful. ", "negative"},
	}
	fmt.Println("Started")
	handleRequests()
}
