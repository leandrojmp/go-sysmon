package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config configuration variables
type Config struct {
	ListenAddress string
	TCPFile       string
}

//Configuration values
var Configuration Config

//LoadConfig load the configuration
func LoadConfig(configFile string) {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("error:", err)
	}
	json.Unmarshal([]byte(content), &Configuration)
}
