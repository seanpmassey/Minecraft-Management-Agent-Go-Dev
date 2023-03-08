package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
	//fmt.Println("GLOBAL=> ", globalconfig)

}

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("minecraftagent: ")
	log.SetFlags(0)

	serverstring := globalconfig.Server + ":" + globalconfig.Port

	userCMD := flag.NewFlagSet("User", flag.ExitOnError)
	listusers := userCMD.Bool("listusers", false, "listusers")
	//userName := userCMD.String("username", "", "username")

	opsCMD := flag.NewFlagSet("Ops", flag.ExitOnError)
	addops := opsCMD.Bool("addops", false, "addops")
	removeops := opsCMD.Bool("removeops", false, "removeops")
	opsuser := opsCMD.String("opsuser", "", "opsuser")

	serverCMD := flag.NewFlagSet("Server", flag.ExitOnError)
	saveall := serverCMD.Bool("saveall", false, "saveall")
	//userName := userCMD.String("username", "", "username")

	fmt.Println(os.Args)
	switch os.Args[1] {
	case "user":
		userCMD.Parse(os.Args[2:])
		//listusers := *listusers
		//var listusersbool = listusers

		if *listusers == true {
			//fmt.Println("listusers is true")
			userlist(serverstring, globalconfig.Password)
		}
	case "ops":
		opsCMD.Parse(os.Args[2:])

		if *addops == true {
			opsadd(serverstring, globalconfig.Password, *opsuser)
		}
		if *removeops == true {
			opsremove(serverstring, globalconfig.Password, *opsuser)
		}
	case "server":
		serverCMD.Parse(os.Args[2:])

		if *saveall == true {
			serversaveall(serverstring, globalconfig.Password)
		}
	default:
		fmt.Println("expected 'user' or 'ops' subcommands")
		os.Exit(1)

	}

	//userlist(serverstring, globalconfig.Password)
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
