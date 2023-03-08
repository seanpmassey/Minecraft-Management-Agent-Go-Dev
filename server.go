package main

import (
	"log"
	//"strings"
)

func serversaveall(serverstring string, password string) {

	command = "save-all"
	response, err := rconconnect(serverstring, password, command)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
