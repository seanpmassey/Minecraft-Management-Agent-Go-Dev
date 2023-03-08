package main

import (
	"log"
	//"strings"
)

func opsadd(serverstring string, password string, opsuser string) {

	command = "ops" + opsuser
	response, err := rconconnect(serverstring, password, command)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)

	//fmt.Println(response)
	//userlist := parseusers[1:]
	//fmt.Println(userlist)
}

func opsremove(serverstring string, password string, opsuser string) {

	command = "deop" + opsuser
	response, err := rconconnect(serverstring, password, command)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)

	//fmt.Println(response)
	//userlist := parseusers[1:]
	//fmt.Println(userlist)
}
