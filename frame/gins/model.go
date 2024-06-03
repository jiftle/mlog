package gins

type Cnf struct {
	Logger Logger `json:"logger"`
}
type Logger struct {
	Path   string `json:"path"`
	Level  string `json:"level"`
	Stdout bool   `json:"stdout"`
}
