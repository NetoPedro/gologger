package gologger

const (
	CONSOLE       string = "console"
	FILE          string = "file"
	ELASTICSEARCH string = "es"
	SimpleLog     string = "simple"
	ColoredLog    string = "color"
)

type GoLogger struct {
	Logger
}

func GetLogger(selector ...string) GoLogger {
	if len(selector) == 0 {
		return GoLogger{Logger{CONSOLE, ColoredLog}}
	}
	return GoLogger{Logger{selector[0], selector[1]}}
}

func (log GoLogger) Log(message interface{}) {
	logPrinter(LogInstance{LogType: "LOG", Body: message, LoggerInit: log.Logger})
}

func (log GoLogger) Message(message interface{}) {
	logPrinter(LogInstance{LogType: "MSG", Body: message, LoggerInit: log.Logger})
}

func (log GoLogger) Info(message interface{}) {
	logPrinter(LogInstance{LogType: "INF", Body: message, LoggerInit: log.Logger})
}

func (log GoLogger) Warn(message interface{}) {
	logPrinter(LogInstance{LogType: "WRN", Body: message, LoggerInit: log.Logger})
}

func (log GoLogger) Debug(message interface{}) {
	logPrinter(LogInstance{LogType: "DBG", Body: message, LoggerInit: log.Logger})
}

func (log GoLogger) Error(message interface{}) {
	logPrinter(LogInstance{LogType: "ERR", Body: message, LoggerInit: log.Logger})
}

func (log GoLogger) Fatal(message interface{}) {
	logPrinter(LogInstance{LogType: "CRT", Body: message, LoggerInit: log.Logger})
}

func (log GoLogger) ReplaceMessage(message interface{}) {
	logPrinter(LogInstance{LogType: "RSS", Body: message, LoggerInit: log.Logger})
}
