package library

import (
	"encoding/json"
	"io/ioutil"
)

// config struct
type Config struct {
	CurrentEnv string      `json:"current_env"`
	Env        interface{} `json:"env"`
}

// ReadFromJson read the Config from a JSON file.
func ReadFromJson(path string) *Config {
	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var cfg *Config

	err = json.Unmarshal(jsonByte, &cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}

func GetConfigPath() string {
	return "./config/config.json"
}
