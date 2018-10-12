// logger project logger.go
package logger

import (
	errs "errorshandler"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"time"

	logrus "github.com/sirupsen/logrus"
)

var logLevel int = 2
var levels map[string]int

type KsiLogger struct {
	logger   *logrus.Logger
	loglevel int
	prefix   string
}

var logger = new(KsiLogger)
var isInited bool = false

func (log *KsiLogger) Output(v ...interface{}) {
	file, line := caller(3)
	now := time.Now()

	fmt.Print(now.Format("01-02-2006 15:04:05"), " ", path.Base(file)+":", line, " ", fmt.Sprintln(v...))
}

func createLogLevels() {
	if len(levels) == 0 {
		levels = make(map[string]int)
		levels["debug"] = 1
		levels["info"] = 2
		levels["error"] = 3
		logger.logger = logrus.New()
		isInited = true
	}
}

func SetLogLevel(level string) {
	createLogLevels()
	if theLevel, ok := levels[level]; ok {

		logger.loglevel = theLevel
		logrusLevel, err := logrus.ParseLevel(level)
		if err == nil {
			logLevel = theLevel
			logger.logger.Level = logrusLevel
		}
		_, file, _, _ := runtime.Caller(1)
		fmt.Println(file)

	}
}

func Info(v ...interface{}) {
	if isLevelEnabled("info") {
		logger.Output(v...)
	}
}

func Debug(v ...interface{}) {
	if isLevelEnabled("debug") {
		logger.Output(v...)
	}
}

func Error(v ...interface{}) {
	if isLevelEnabled("error") {
		logger.Output(v...)
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
	}
}

func isLevelEnabled(level string) bool {
	if !isInited {
		createLogLevels()
	}

	return levels[level] >= logger.loglevel
}
