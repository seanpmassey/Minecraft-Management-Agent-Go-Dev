package main

import (
	"fmt"
	"log"
	"strings"
)

func listusers(serverstring string, password string) {

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
