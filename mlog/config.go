package mlog

import "github.com/BurntSushi/toml"

type Config struct {
	Path        string
	Level       string
	Stdout      bool
	stackDeepth int
}

func loadConfig() (cnf *Config, err error) {
	m := make(map[string]interface{})
	file := "config/config.toml"

	_, err = toml.DecodeFile(file, &m)
	if err != nil {
		return
	}

	m = m["logger"].(map[string]interface{})
	cnf = &Config{
		stackDeepth: 3,
		Path:        m["path"].(string),
		Level:       m["level"].(string),
		Stdout:      m["stdout"].(bool),
	}

	return
}
