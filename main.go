package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person is a Subscriber
type Person struct {
	ID         string      `json:"id,omitempty"`
	Username   string      `json:"username,omitempty"`
	Subscribed *Subscribed `json:"subs,omitempty"`
}

// Subscribed is list of Topics
type Subscribed struct {
	Topic string `json:"topics,omitempty"`
}

// Topic list
type Topic struct {
	ID    string   `json:"id,omitempty"`
	Topic string   `json:"topic,omitempty"`
	News  []string `json:"news,omitempty"`
}

var subscriber []Person
var newslist []Topic

func init() {
	x.AddNews("Football", "some news")
	x.AddNews("Football", "second news")
	x.AddNews("Tenis", "tenis news")
	subscriber = append(subscriber, Person{ID: "1", Username: "Stephen", Subscribed: &Subscribed{Topic: "Football - Tenis"}})
	subscriber = append(subscriber, Person{ID: "2", Username: "John", Subscribed: &Subscribed{Topic: "Football"}})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/register", Register).Methods("POST")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{username}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/subscribe/{username}", GetNews).Methods("GET")
	router.HandleFunc("/subscribe/{username}/{topic}", Subscribe).Methods("GET")
	router.HandleFunc("/subscribed/{username}/topics", UserTopics).Methods("GET")
	router.HandleFunc("/publish/{topic}/{news}", Publish).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
