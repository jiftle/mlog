package mlog

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

func (l *Logger) GetStack() (out string) {
	index := l.config.stackDeepth
	depth := index - l.config.stackDeepth + 1
	pc, file, line, ok := runtime.Caller(index)
	if !ok {
		return
	}
	fn := runtime.FuncForPC(pc)

	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteString("Stack:\n")

	slineNum := fmt.Sprintf("%d. ", depth)
	buffer.WriteString(fmt.Sprintf("%s %s\n", slineNum, fn.Name()))
	buffer.WriteString(fmt.Sprintf("%s %s:%d", strings.Repeat(" ", len(slineNum)), file, line))
	buffer.WriteString("\n")
	out = buffer.String()
	return
}
