package logger

import "fmt"

func (l *Logger) LogLines(lines [][]rune) {
	for k, v := range lines {
		l.Log(fmt.Sprintf("Index %d, line: %v", k, v))
	}
}
