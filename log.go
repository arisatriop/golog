package golog

import (
	"os"

	"github.com/arisatriop/golog/help"
	"github.com/sirupsen/logrus"
)

type Logging struct {
	Activity string
	Level    string
	Message  string
	Location string
	Time     string
	User     string
}

func initLog() *logrus.Logger {
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	return logger
}

func initJsonLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	return logger
}

func LogStart(activity string, user string) {
	logger := initLog()
	logger.Println()
	logger.Printf("A new proses '%s' by '%s' is starting.", activity, user)
}

func LogEnd(activity string, user string) {
	logger := initLog()
	logger.Printf("'%s' by '%s' perfectly complete.", activity, user)
}

func LogError(activity, loc, msg, user string) {
	time := help.GetTime().Format("2006-01-02 15:04:05")

	logger := initJsonLog()
	logger.WithFields(logrus.Fields{"activity": activity, "location": loc, "local_time": time, "user": user}).Error(msg)
}

// func LogInfo(activity, user string) {
// 	// time := help.GetTime().Format("2006-01-02 15:04:05")

// 	logger := initLog()
// 	logger.Info()
// }
