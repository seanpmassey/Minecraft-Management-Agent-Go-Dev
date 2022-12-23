package main

import (
	"fmt"
	"log"
	"strings"

	mcrcon "github.com/Kelwing/mc-rcon"
)

// declare global variables
var server string = "minecraft-s1.lan.seanmassey.net"
var port string = "25575"
var password string = "VMware1!"
var command string

func rconconnect(server string, password string, command string) (string, error) {
	conn := new(mcrcon.MCConn)
	err := conn.Open(server, password)
	if err != nil {
		log.Fatalln("Open failed", err)
	}
	defer conn.Close()

	err = conn.Authenticate()
	if err != nil {
		log.Fatalln("Auth failed", err)
	}

	response := ""
	response, err = conn.SendCommand(command)
	if err != nil {
		log.Fatalln("Command failed", err)
	}

	return response, err

}

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("minecraftagent: ")
	log.SetFlags(0)

	// Declare Server Variable Values
	//server := "minecraft-s1.lan.seanmassey.net"
	//port := "25575"
	//password := "VMware1!"
	//command := "list"

	serverstring := server + ":" + port

	listusers(serverstring)
}

func listusers(serverstring string) {

	command = "list"
	response, err := rconconnect(serverstring, password, command)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)

	parseresponse := strings.Split(response, ":")
	trimusers := strings.TrimSpace(parseresponse[1])
	parseusers := strings.Split(trimusers, ",")
	fmt.Println(parseusers)
	//userlist := parseusers[1:]
	//fmt.Println(userlist)
}
