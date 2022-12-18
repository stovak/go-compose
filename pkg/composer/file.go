package composer

import (
	"encoding/json"
	"log"
	"os"
)

type File struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Version     string `mapstructure:"version"`
}

func New(filename string) File {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload File
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Let's print the unmarshalled data!
	log.Printf("payload: %+v", payload)
	return payload
}
