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

	flag.StringVar(&username, "username", "", "your username")
	flag.StringVar(&password, "password", "", "your password")
	flag.StringVar(&roomID, "roomid", "", "the room id you want to message")
	flag.StringVar(&message, "message", "", "the message you want to convey")

	flag.Parse()

	for _, str := range []string{username, password, roomID, message} {
		if len(str) == 0 {
			flag.Usage()
			return
		}
	}

	client := matrix.New(username, password)

	if err := client.Send(roomID, message); err != nil {
		log.Fatal(err)
	}
}
