package gologger

type Logger struct {
	PrinterType string
	Location    string
	MinLevel	int
}

type LogInstance struct {
	LogType    string
	Body    interface{}
	LoggerInit Logger
}
