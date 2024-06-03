package gins

import (
	"encoding/json"
	"os"
)

func loadConfig() (m map[string]interface{}, err error) {
	m = make(map[string]interface{})

	file, err := os.Open("config/config.json")
	if err != nil {
		return
	}
	defer file.Close()

	cnf := Cnf{}
	bytFile := make([]byte, 1024)
	n, err := file.Read(bytFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytFile[:n], &cnf)
	if err != nil {
		return
	}

	m["path"] = cnf.Logger.Path
	m["level"] = cnf.Logger.Level
	m["stdout"] = cnf.Logger.Stdout
	return
}
