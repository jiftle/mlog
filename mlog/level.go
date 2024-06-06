package mlog

import (
	"fmt"
	"strings"
)

const (
	LEVEL_DEBU = 1 << iota
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERRO
)

var levelStringMap = map[string]int{
	"DEBU":  LEVEL_DEBU | LEVEL_INFO | LEVEL_WARN | LEVEL_ERRO,
	"DEBUG": LEVEL_DEBU | LEVEL_INFO | LEVEL_WARN | LEVEL_ERRO,
	"INFO":  LEVEL_INFO | LEVEL_WARN | LEVEL_ERRO,
	"WARN":  LEVEL_WARN | LEVEL_ERRO,
	"ERRO":  LEVEL_ERRO,
	"ERROR": LEVEL_ERRO,
}

func GetLevelByStr(level string) int {
	return levelStringMap[level]
}

func (s *Logger) SetLevelStr(levelStr string) error {
	level, ok := levelStringMap[strings.ToUpper(levelStr)]
	if !ok {
		return fmt.Errorf("not support level")
	}
	s.config.level = level
	return nil
}

func (s *Logger) checkLevel(level int) bool {
	return s.config.level&level > 0
}
