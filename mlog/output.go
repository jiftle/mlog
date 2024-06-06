package mlog

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func (s *Logger) printToStdout(level string, format string, v ...interface{}) {
	buffer := bytes.NewBuffer([]byte{})

	// head time
	stime := time.Now().Format("2006-01-02 15:04:05.000")
	buffer.WriteString(stime)
	buffer.WriteString(" ")

	// head level
	slevel := "[ERRO]"
	switch level {
	case "DEBU":
		slevel = "[DEBU]"
		slevel = color.YellowString("%s", slevel)
	case "INFO":
		slevel = "[INFO]"
		slevel = color.GreenString("%s", slevel)
	case "WARN":
		slevel = "[WARN]"
		slevel = color.MagentaString("%s", slevel)
	case "ERRO":
		slevel = "[ERRO]"
		slevel = color.RedString("%s", slevel)
	default:
		slevel = "[DEBU]"
		slevel = color.RedString("%s", slevel)
	}
	buffer.WriteString(slevel)
	buffer.WriteString(" ")

	values := v
	tempStr := ""
	if len(format) > 0 {
		tempStr = fmt.Sprintf(format, v...)
	} else {
		for _, v := range values {
			tempStr = ConvAnyToString(v)
		}
	}
	buffer.WriteString(tempStr)

	buffer.WriteString("\n")
	if strings.EqualFold(level, "ERRO") {
		buffer.WriteString(s.GetStack())
	}

	fmt.Fprint(color.Output, buffer.String())
}

func (s *Logger) printToFile(level string, format string, v ...interface{}) {
	buffer := bytes.NewBuffer([]byte{})

	// head time
	stime := time.Now().Format("2006-01-02 15:04:05.000")
	buffer.WriteString(stime)
	buffer.WriteString(" ")

	// head level
	slevel := "[ERRO]"
	switch level {
	case "DEBU":
		slevel = "[DEBU]"
	case "INFO":
		slevel = "[INFO]"
	case "WARN":
		slevel = "[WARN]"
	case "ERRO":
		slevel = "[ERRO]"
	default:
		slevel = "[DEBU]"
	}
	buffer.WriteString(slevel)
	buffer.WriteString(" ")

	values := v
	tempStr := ""
	if len(format) > 0 {
		tempStr = fmt.Sprintf(format, v...)
	} else {
		for _, v := range values {
			tempStr = ConvAnyToString(v)
		}
	}
	buffer.WriteString(tempStr)

	buffer.WriteString("\n")
	if strings.EqualFold(level, "ERRO") {
		buffer.WriteString(s.GetStack())
	}

	file := fmt.Sprintf("%s/%s.log", s.config.Path, time.Now().Format("2006-01-02"))
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	f.Write(buffer.Bytes())
}
