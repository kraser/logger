// logger project logger.go
package logger

import (
	"fmt"
	//"log"
	"runtime"
)

var logLevel int = 2
var levels map[string]int

func createLogLevels() {
	if len(levels) == 0 {
		levels = make(map[string]int)
		levels["debug"] = 1
		levels["info"] = 2
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

func caller(calldepth int) (string, int) {
	var ok bool
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	return file, line
}
