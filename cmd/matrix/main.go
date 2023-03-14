package main

import (
	"flag"
	"log"

	"github.com/prophittcorey/matrix"
)

func main() {
	var username string
	var password string
	var roomID string
	var message string
	var subject string

	flag.StringVar(&username, "username", "", "your username")
	flag.StringVar(&password, "password", "", "your password")
	flag.StringVar(&roomID, "roomid", "", "the room id you want to message")
	flag.StringVar(&subject, "subject", "", "the subject or notification message")
	flag.StringVar(&message, "message", "", "the message you want to convey")

	flag.Parse()

	for _, str := range []string{username, password, roomID, message} {
		if len(str) == 0 {
			flag.Usage()
			return
		}
	}

	client := matrix.New(username, password)

	if err := client.Send(roomID, subject, message); err != nil {
		log.Fatal(err)
	}
}
