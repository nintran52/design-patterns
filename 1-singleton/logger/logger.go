package logger

var loggerInstance *Logger

type Logger struct {
	Level string
}

func GetLogger() *Logger {
	if loggerInstance == nil {
		loggerInstance = &Logger{Level: "INFO"}
	}
	return loggerInstance
}
