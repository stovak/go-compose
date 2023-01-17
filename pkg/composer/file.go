package composer

import (
	"log"
	"os"

	"encoding/json"
)

type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Homepage string `json:"homepage"`
	Role     string `json:"role"`
}

type FundingLink struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Repository struct{}

type AutoloadSpec struct {
	Psr4                map[string]string `json:"psr-4"`
	Psr0                map[string]string `json:"psr-0"`
	Classmap            []string          `json:"classmap"`
	Files               []string          `json:"files"`
	ExcludeFromClassmap []string          `json:"exclude-from-classmap"`
	OptimizeAutoloader  bool              `json:"optimize-autoloader"`
}

type ConfigSpec struct{}

type File struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Version     string            `json:"version"`
	Type        string            `json:"type"`
	Keywords    []string          `json:"keywords"`
	Homepage    string            `json:"homepage"`
	Readme      string            `json:"readme"`
	Time        string            `json:"time"`
	License     string            `json:"license"`
	Authors     []Author          `json:"authors"`
	Support     map[string]string `json:"support"`
	Funding     []FundingLink     `json:"funding"`
	Bin         string            `json:"bin"`

	MinimumStability string   `json:"minimum-stability"`
	PreferStable     bool     `json:"prefer-stable"`
	IncludePath      []string `json:"include-path"`
	TargetDir        string   `json:"target-dir"`

	Require    map[string]string `json:"require"`
	RequireDev map[string]string `json:"require-dev"`
	Conflict   map[string]string `json:"conflict"`
	Replace    map[string]string `json:"replace"`
	Provide    map[string]string `json:"provide"`
	Suggest    map[string]string `json:"suggest"`

	Autoload    AutoloadSpec `json:"autoload"`
	AutoloadDev AutoloadSpec `json:"autoload-dev"`

	Config    map[string]interface{} `json:"config"`
	Extra     map[string]interface{} `json:"extra"`
	Scripts   map[string]interface{} `json:"scripts"`
	Abandoned string                 `json:"abandoned"`
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
	return payload
}
