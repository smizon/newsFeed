package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GetUsers - displays all subscribers
func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(subscriber)
}

// Register - creates new subscriber
func Register(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	subscriber = append(subscriber, person)
	json.NewEncoder(w).Encode(subscriber)
}

// DeleteUser - deletes subscriber from users
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range subscriber {
		if item.Username == params["username"] {
			subscriber = append(subscriber[:index], subscriber[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(subscriber)
	}
}

// clean - cleans list string
func clean(d string) string {
	t := strings.Trim(d, "&{ }")
	return t
}

// Subscribe to Topic
func Subscribe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range subscriber {
		if item.Username == params["username"] {
			// formalate Topic List / clean
			tl := clean(fmt.Sprintf("%s ", item.Subscribed))
			if strings.Contains(tl, params["topic"]) {
				return
			}
			topicList := tl + " - " + params["topic"]
			tList := &Subscribed{Topic: topicList}
			message := "Subscribed to " + params["topic"]
			// Save
			subscriber[index].Subscribed = tList
			json.NewEncoder(w).Encode(message)
			return
		}
	}
}

// UserTopics - return subscribed topics list
func UserTopics(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range subscriber {
		if item.Username == params["username"] {
			tl := clean(fmt.Sprintf("%s ", item.Subscribed))
			tlist := strings.Split(tl, " - ")
			fmt.Println(tlist)
			json.NewEncoder(w).Encode(tlist)
			return
		}
	}
}

// GetNews - displays news topics for Subscriber (needs to display topic news)
func GetNews(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	t := ""
	i := []string{}

	for _, item := range subscriber {
		if item.Username == params["username"] {
			tl := clean(fmt.Sprintf("%s ", item.Subscribed))
			tlist := strings.Split(tl, " - ")

			for _, ns := range tlist {
				val := clean(fmt.Sprintf("%s ", ns))
				t, i = getAllChannel(x, val)
				newslist = append(newslist, Topic{ID: "1", Topic: t, News: i})
			}
			fmt.Println(newslist)
			json.NewEncoder(w).Encode(newslist)
		}
	}
	newslist = nil

}

// Publish News
func Publish(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	topic := params["topic"]
	news := params["news"]
	x.AddNews(topic, news)
	json.NewEncoder(w).Encode("Published")
}
