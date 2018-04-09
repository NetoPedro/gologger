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
		return GoLogger{Logger{CONSOLE, ColoredLog,0}}
	}
	level, _ := strconv.Atoi(selector[2])
	return GoLogger{Logger{selector[0], selector[1],level}}
}

func (log GoLogger) Log(message interface{}) {
	if log.MinLevel <= LOG {
		logPrinter(LogInstance{LogType: "LOG", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) Message(message interface{}) {
	logPrinter(LogInstance{LogType: "MSG", Body: message, LoggerInit: log.Logger})
}

func (log GoLogger) Info(message interface{}) {
	if log.MinLevel <= INFO {
		logPrinter(LogInstance{LogType: "INF", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) Warn(message interface{}) {
	if log.MinLevel <= WARN {
		logPrinter(LogInstance{LogType: "WRN", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) Debug(message interface{}) {
	if log.MinLevel <= DEBUG {
		logPrinter(LogInstance{LogType: "DBG", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) Error(message interface{}) {
	if log.MinLevel <= ERROR {
		logPrinter(LogInstance{LogType: "ERR", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) Fatal(message interface{}) {
	if log.MinLevel <= FATAL {
		logPrinter(LogInstance{LogType: "CRT", Body: message, LoggerInit: log.Logger})
	}
}

func (log GoLogger) ReplaceMessage(message interface{}) {
	logPrinter(LogInstance{LogType: "RSS", Body: message, LoggerInit: log.Logger})
}
