package logging

import (
	"fmt"
	"strings"

	"github.com/xairline/goplane/xplm/utilities"
)

// Information about a loglevel
type Level struct {
	number byte   // number then log levels
	name   string //name then Log levels
}

var (
	// Loglevel for trace messages
	Trace_Level = Level{1, "TRACE"}
	// Loglevel for debug reports
	Debug_Level = Level{2, "DEBUG"}
	// Loglevel for info messages
	Info_Level = Level{3, "INFO"}
	// Loglevel for warnings
	Warning_Level = Level{4, "WARNING"}
	// Loglevel for error
	Error_Level = Level{5, "ERROR"}

	// Level from which the messages are output
	MinLevel = Info_Level
	// Current PluginName
	PluginName = "<unknown>"
)

// If a string determines the corresponding loglevel.Possible values are: trace, debug, info, warning, error.
// If another string is used, then the method returns the info level
func GetLevelFromString(level string) Level {
	switch strings.ToUpper(level) {
	case "TRACE":
		return Trace_Level
	case "DEBUG":
		return Debug_Level
	case "INFO":
		return Info_Level
	case "WARNING":
		return Warning_Level
	case "ERROR":
		return Error_Level
	default:
		return Info_Level
	}
}

// writes a trace message into the log file
func Trace(msg string) {
	writeMessage(Trace_Level, msg)
}

// writes a formatted trace message into the log file
func Tracef(format string, a ...interface{}) {
	if Trace_Level.number >= MinLevel.number {
		Trace(fmt.Sprintf(format, a...))
	}
}

// writes a debug message into the log file
func Debug(msg string) {
	writeMessage(Debug_Level, msg)
}

// writes a formatted debug message into the log file
func Debugf(format string, a ...interface{}) {
	if Debug_Level.number >= MinLevel.number {
		Debug(fmt.Sprintf(format, a...))
	}
}

// writes an info message into the log file
func Info(msg string) {
	writeMessage(Info_Level, msg)
}

// writes a formatted information message into the log file
func Infof(format string, a ...interface{}) {
	if Info_Level.number >= MinLevel.number {
		Info(fmt.Sprintf(format, a...))
	}
}

// writes a warning in the log file
func Warning(msg string) {
	writeMessage(Warning_Level, msg)
}

// writes a formatted warning into the log file
func Warningf(format string, a ...interface{}) {
	if Warning_Level.number >= MinLevel.number {
		Warning(fmt.Sprintf(format, a...))
	}
}

// writes an error message to the log file
func Error(msg string) {
	writeMessage(Error_Level, msg)
}

// writes a formatted error message into the log file
func Errorf(format string, a ...interface{}) {
	if Error_Level.number >= MinLevel.number {
		Error(fmt.Sprintf(format, a...))
	}
}

func writeMessage(level Level, msg string) {
	if level.number >= MinLevel.number {
		utilities.DebugString(fmt.Sprintf("[%v] %v: %v\n", PluginName, level.name, msg))
	}
}
