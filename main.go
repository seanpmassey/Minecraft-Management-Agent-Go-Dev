package main

import (
	"flag"

	//"io/ioutil"
	"log"
	"os"

	//"strings"

	mcrcon "github.com/Kelwing/mc-rcon"
	//yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Server          string `yaml:"server"`
	Port            string `yaml:"port"`
	Password        string `yaml:"password"`
	Properties      string `yaml:"properties"`
	DefaultGameMode string
}

var command string
var globalconfig Config

func init() {
	PopulateConfig("config.yml")

	if globalconfig.Properties != "" {
		ReadServerProperties(globalconfig.Properties)
	}
}

func rconconnect(server string, password string, command string) (string, error) {
	conn := new(mcrcon.MCConn)
	err := conn.Open(server, password)
	if err != nil {
		log.Fatalln("Error: Open failed", err)
	}
	defer conn.Close()

	err = conn.Authenticate()
	if err != nil {
		log.Fatalln("Error: Auth failed", err)
	}

	response := ""
	response, err = conn.SendCommand(command)
	if err != nil {
		log.Fatalln("Error: Command failed", err)
	}

	return response, err

}

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("minecraftagent: ")
	log.SetFlags(0)

	if globalconfig.Server == "" {
		log.Println("Info: No Server Name Set. Defaulting to localhost.")
		globalconfig.Server = "localhost"
	}

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
	setweather := serverCMD.Bool("setweather", false, "setweather")
	weathertype := serverCMD.String("weathertype", "", "weathertype")
	getDefaultgamemode := serverCMD.Bool("getdefaultgamemode", false, "getdefaultgamemode")
	//userName := userCMD.String("username", "", "username")

	//fmt.Println(os.Args)
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
		if *setweather == true {
			switch *weathertype {
			case "clear", "rain", "thunder":
				setserverweather(serverstring, globalconfig.Password, *weathertype)
			default:
				log.Println("Unknown Weather Type. Please select Clear, Rain, or Thunder.")
			}
		}
		if *getDefaultgamemode == true {
			//filepath := "server.properties"
			getdefaultgamemode(globalconfig.Properties)
		}
	default:
		log.Println("expected 'user' 'ops' or 'server' subcommands")
		os.Exit(1)

	}

}
