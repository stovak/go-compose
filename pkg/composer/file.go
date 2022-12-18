package composer

import (
	"log"
	"os"

	"encoding/json"
)

type Author struct {
	Name     string `mapstructure:"name"`
	Email    string `mapstructure:"email"`
	Homepage string `mapstructure:"homepage"`
	Role     string `mapstructure:"role"`
}

type File struct {
	Name        string   `mapstructure:"name"`
	Description string   `mapstructure:"description"`
	Version     string   `mapstructure:"version"`
	Type        string   `mapstructure:"type"`
	Keywords    []string `mapstructure:"keywords"`
	Homepage    string   `mapstrucutre:"homepage"`
	Readme      string   `mapstructure:"readme"`
	Time        string   `mapstructure:"time"`
	License     string   `mapstructure:"license"`
	Authors     []Author `mapstructure:"authors"`
}

func New(filename string) File {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshal the data into `payload`
	var payload File
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Let's print the unmarshalled data!
	log.Printf("payload: %+v", payload)
	return payload
}
