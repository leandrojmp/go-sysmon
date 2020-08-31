package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config configuration variables
type Config struct {
	ListenAddress string
	TCPFile       string
}

//WarnLogger warning logs
//InfoLogger info logs
//ErrLogger warning logs
var (
	WarnLogger *log.Logger
	InfoLogger *log.Logger
	ErrLogger  *log.Logger
)

//CreateLoggers cria loggers
func CreateLoggers() {
	InfoLogger = log.New(os.Stderr, "INFO: ", log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	WarnLogger = log.New(os.Stderr, "WARNING: ", log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	ErrLogger = log.New(os.Stderr, "ERROR: ", log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
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
