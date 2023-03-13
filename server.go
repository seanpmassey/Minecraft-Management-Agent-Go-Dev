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

func setserverweather(serverstring string, password string, weathertype string) {
	command = "weather " + weathertype
	response, err := rconconnect(serverstring, password, command)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}

func setdefaultgamemode(serverstring string, password string, gamemode string) {

}

func getdefaultgamemode(filepath string) {
	ReadPropertiesDefaultGameMode(filepath)
}

func setgametime(serverstring string, password string) {

}
