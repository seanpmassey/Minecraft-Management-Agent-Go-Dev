package main

import (
	"fmt"
	"io/ioutil"
	"log"

	//"strings"

	mcrcon "github.com/Kelwing/mc-rcon"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Server   string `yaml:"server"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

var command string
var globalconfig Config

func init() {
	PopulateConfig("config.yml")
}

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

func PopulateConfig(filePath string) {
	data, err := ioutil.ReadFile(filePath)

	// Read the file
	//data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(data)

	err = yaml.Unmarshal(data, &globalconfig)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("GLOBAL=> ", globalconfig)

}

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("minecraftagent: ")
	log.SetFlags(0)

	serverstring := globalconfig.Server + ":" + globalconfig.Port

	listusers(serverstring, globalconfig.Password)
}

// func listusers(serverstring string) {

// 	command = "list"
// 	response, err := rconconnect(serverstring, password, command)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(response)

// 	parseresponse := strings.Split(response, ":")
// 	trimusers := strings.TrimSpace(parseresponse[1])
// 	parseusers := strings.Split(trimusers, ",")
// 	fmt.Println(parseusers)
// 	//userlist := parseusers[1:]
// 	//fmt.Println(userlist)
// }
