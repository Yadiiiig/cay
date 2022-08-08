package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	File *os.File
}

func Create(logs string) (Logger, error) {
	log_file, err := os.OpenFile(logs, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return Logger{}, err
	}

	return Logger{File: log_file}, nil
}

func (l *Logger) Clean() {
	l.File.Truncate(0)
	l.File.Seek(0, 0)
	l.File.Sync()
}

func (l *Logger) Log(line string) {
	h, m, s := time.Now().Clock()
	if _, err := l.File.WriteString(fmt.Sprintf("%d:%d:%d - %s\n", h, m, s, line)); err != nil {
		fmt.Println("Logger failed")
	}

	l.File.Sync()
}
