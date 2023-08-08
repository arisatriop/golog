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

func LogStart(params ...string) {
	logger := initLog()
	if len(params) == 1 {
		logger.Printf("A new process '%s' is started.", params[0])
		return
	}

	if len(params) == 2 {
		logger.Printf("A new process '%s' by '%s' is started.", params[0], params[1])
		return
	}

	logger.Println("A new process is started.")
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

func LogInfo(activity, user string) {
	// time := help.GetTime().Format("2006-01-02 15:04:05")

	logger := initLog()
	logger.Info()
}
