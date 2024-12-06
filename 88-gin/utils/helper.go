package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

var (
	LogChan chan string
)

func init() {
	if LogChan == nil {
		LogChan = make(chan string)
	}
	go ProcessLogs()
}

func ProcessLogs() {
	for line := range LogChan {
		println(line)
	}
}

func GetLogHeader(str string) string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"), " ", filepath.Base(file), " ", line, str, " ")
}
