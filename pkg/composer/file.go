package composer

type File struct {
	name string `mapstructure:"name"`
	description string `mapstructure:"description"`
    version string `mapstructure:"version"`
}
