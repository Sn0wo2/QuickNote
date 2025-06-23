package config

type Config struct {
	Listener struct {
		Address string `yaml:"address"`
		Static  string `optional:"true" yaml:"static"`
	} `yaml:"listener"`
	Logger struct {
		Level string `yaml:"level"`
	} `yaml:"logger"`
	Database struct {
		Type string `yaml:"type"`
		URL  string `yaml:"url"`
	} `yaml:"database"`
}
