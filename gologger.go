package gologger

import "strconv"

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
const DEBUG = 0
const LOG = 1
const INFO = 2
const WARN = 3
const ERROR = 4
const FATAL = 5

func GetLogger(selector ...string) GoLogger {
	if len(selector) == 0 {
		return GoLogger{Logger{CONSOLE, ColoredLog}}
	}
	return GoLogger{Logger{selector[0], selector[1]}}
}

func (log GoLogger) Log(message interface{}, minLevel int) {
	if minLevel <= LOG {
		logPrinter(LogInstance{LogType: "LOG", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) Message(message interface{}) {
	logPrinter(LogInstance{LogType: "MSG", Body: message, LoggerInit: log.Logger})
}

func (log GoLogger) Info(message interface{}, minLevel int) {
	if minLevel <= INFO {
		logPrinter(LogInstance{LogType: "INF", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) Warn(message interface{}, minLevel int) {
	if minLevel <= WARN {
		logPrinter(LogInstance{LogType: "WRN", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) Debug(message interface{}, minLevel int) {
	if minLevel <= DEBUG {
		logPrinter(LogInstance{LogType: "DBG", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) Error(message interface{}, minLevel int) {
	if minLevel <= ERROR {
		logPrinter(LogInstance{LogType: "ERR", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) Fatal(message interface{}, minLevel int) {
	if minLevel <= FATAL {
		logPrinter(LogInstance{LogType: "CRT", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) ReplaceMessage(message interface{}) {
	logPrinter(LogInstance{LogType: "RSS", Body: message, LoggerInit: log.Logger})
}
