package rlog

import (
	"fmt"
	"time"
)

type Logger struct {
	Mode string
}

func (l *Logger) getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (l *Logger) Error(err string) {
	t := l.getTime()
	fmt.Println(fmt.Sprintf("%s %s %s", t, "ERROR:", err))
}

func (l *Logger) Info(info string) {
	if l.Mode == "info" {
		t := l.getTime()
		fmt.Println(fmt.Sprintf("%s %s %s", t, "INFO:", info))
	}

}

func (l *Logger) Debug(debug string) {
	if l.Mode == "debug" {
		t := l.getTime()
		fmt.Println(fmt.Sprintf("%s %s %s", t, "DEBUG:", debug))

	}
}

func NewLogger(mode string) *Logger {
	return &Logger{Mode: mode}
}
