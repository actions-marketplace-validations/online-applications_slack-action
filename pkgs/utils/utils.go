package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"slack-action/pkgs/slack"
)

func GetEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}


func GetJsonValue(key, file string) (string, error){
	// Generic func to get value of key from JSON file
    // Open our jsonFile
    jsonFile, err := os.Open(file)
    if err != nil {
        log.Println("Got error while trying to open file:", file)
    }
    log.Println("Successfully Opened:", file)
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]string

	if err := json.Unmarshal([]byte(byteValue), &result); err != nil {
        log.Println("Error was found when unmarshaling JSON file:", file)
    }
	
	log.Printf("Searching key: %s inside the file: %s", key, file)
    v  := result[key]
	log.Println("Found key:", v)

	return v, err
}

func GetCliArg(argNum int) string {
	args := os.Args
	if len(args) == argNum {
		return ""
	}
	arg := os.Args[argNum]
	log.Println("Found arg:", arg)
	return arg

}

func ReadFile(filePath string) ([]byte, error) {
	payloadJSON, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Error was found while reading from custom payload file!", err)
	}
	return payloadJSON, err
}

func JsonMarshal(payload slack.Message) ([]byte, error){
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error was found when parsing JSON:", err)
	}
	return payloadJSON, err
}