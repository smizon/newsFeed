package main

import (
	"fmt"
	"strings"
)

type newsChannel map[string][]string

var x = make(newsChannel)

// Saves the news to the Channel
func (s newsChannel) AddNews(topic, value string) {
	_, ok := s[topic]
	if !ok {
		s[topic] = make([]string, 0, 20)
	}
	s[topic] = append(s[topic], value)
}

// Gets the news from the Channel
func (s newsChannel) GetNews(topic string, index int) (string, bool) {
	slice, ok := s[topic]

	if !ok || len(slice) == 0 {
		return "", false
	}
	return s[topic][index], true
}

// Count News in Channel
func (n newsChannel) countNews(topic string) int {
	return len(n[topic])
}

// Get All news in Channel
func getAllChannel(x newsChannel, topic string) (string, []string) {
	news := x.countNews(topic)
	vlist := ""
	for i := 0; i < news; i++ {
		v, ok := x.GetNews(topic, i)
		if ok {
			fmt.Println(v)
			if i == 0 {
				vlist = vlist + v
			}
			if i != 0 {
				vlist = vlist + " - " + v
			}
		} else {
			fmt.Println(`unable to read value for key "key"`)
		}
		// return v
	}
	vl := strings.Split(vlist, " - ")
	return topic, vl
}
