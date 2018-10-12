// logger project logger.go
package logger

import (
	errs "errorshandler"
	"fmt"
	"os"
	"regexp"
	"runtime"
)

var logLevel int = 2
var levels map[string]int

func createLogLevels() {
	if len(levels) == 0 {
		levels = make(map[string]int)
		levels["debug"] = 1
		levels["info"] = 2
		levels["error"] = 3
	}
}

func SetLogLevel(level string) {
	createLogLevels()
	if theLevel, ok := levels[level]; ok {
		logLevel = theLevel
	}
}

func Info(v ...interface{}) {
	createLogLevels()
	if levels["info"] >= logLevel {
		file, line := caller(2)
		fmt.Print(file+":", line, " ", fmt.Sprintln(v...))
	}
}

func Debug(v ...interface{}) {
	createLogLevels()
	if levels["debug"] >= logLevel {
		file, line := caller(2)
		fmt.Print(file+":", line, " ", fmt.Sprintln(v...))
	}
}

func Error(v ...interface{}) {
	createLogLevels()
	if levels["error"] >= logLevel {
		file, line := caller(2)
		fmt.Print(file+":", line, " ", fmt.Sprintln(v...))
	}
}

func caller(calldepth int) (string, int) {
	var ok bool
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	return file, line
}

func CheckHtml(rawUrl string, html string, level string) {
	createLogLevels()

	if levels[level] >= logLevel {
		re := regexp.MustCompile("[^a-zA-Z0-9]+")
		fileName := "/home/robot/" + re.ReplaceAllString(rawUrl, "_")
		fileHandler, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		errs.ErrorHandle(err)
		defer fileHandler.Close()
		fileHandler.Truncate(0)
		fileHandler.WriteString(html)
		Debug(len(html))
	}
}
