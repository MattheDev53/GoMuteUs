package config

import (
	"encoding/json"
	"github.com/charmbracelet/log"
	"os"
)

type Config struct {
	GuildID string
	Token   string
}

var conf, err = os.UserConfigDir()
var ConfigDir = conf + "/GoMuteUs/"
var ConfigFile = ConfigDir + "config.json"

func ReadConfig() Config {
	byteString, err := os.ReadFile(ConfigFile)
	if err != nil {
		log.Fatalf("Error Reading Config File: %v\n", err)
	}
	c := FromJson(byteString)
	return c
}

func FromJson(b []byte) Config {
	var c Config
	err := json.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
