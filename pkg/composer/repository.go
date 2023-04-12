package composer

type Repository struct {
	Type    string            `json:"type"`
	URL     string            `json:"url"`
	Options map[string]string `json:"options"`
	NoDev   bool              `json:"no-dev"`
	OnlyDev bool              `json:"only-dev"`
	Package *Package          `json:"package"`
}

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Dist    struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	}
	Source struct {
		Type      string `json:"type"`
		URL       string `json:"url"`
		Reference string `json:"reference"`
	}
}
