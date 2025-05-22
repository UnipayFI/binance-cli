package config

import "os"

var Config struct {
	APIKey    string
	APISecret string
}

func init() {
	Config.APIKey = os.Getenv("API_KEY")
	Config.APISecret = os.Getenv("API_SECRET")
}
