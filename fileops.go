package main

import (
	//"fmt"
	"io/ioutil"
	"strings"

	"bufio"
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func PopulateConfig(filePath string) {
	data, err := ioutil.ReadFile(filePath)

	// Read the file
	//data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal("error:", err)
		return
	}

	//fmt.Println(data)

	err = yaml.Unmarshal(data, &globalconfig)
	if err != nil {
		log.Fatal("error:", err)
	}
	//fmt.Println("GLOBAL=> ", globalconfig)

}

func ReadServerProperties(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("error:", err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		//fmt.Printf("line: %s\n", scanner.Text())
		text := scanner.Text()

		switch {
		case strings.Contains(text, "rcon.port"):
			rconport := strings.Split(text, "=")
			//fmt.Println(rconport[1])
			globalconfig.Port = rconport[1]
			log.Println("The rcon port is " + globalconfig.Port)
		case strings.Contains(text, "rcon.password"):
			rconpassword := strings.Split(text, "=")
			rconpass := strings.Replace(rconpassword[1], "\\", "", -1)
			//fmt.Println(rconpass)
			globalconfig.Password = rconpass
			log.Println("The rcon password is " + globalconfig.Password)
		case strings.HasPrefix(text, "gamemode"):
			gamemode := strings.Split(text, "=")
			globalconfig.DefaultGameMode = gamemode[1]
			log.Println("The Default Gamemode is: " + globalconfig.DefaultGameMode)
		default:

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func ReadPropertiesDefaultGameMode(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("error:", err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		//fmt.Printf("line: %s\n", scanner.Text())
		text := scanner.Text()

		if strings.HasPrefix(text, "gamemode") {
			gamemode := strings.Split(text, "=")
			globalconfig.DefaultGameMode = gamemode[1]
			log.Println("The Default Gamemode is: " + globalconfig.DefaultGameMode)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
