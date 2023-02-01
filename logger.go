package simplelog

import "os"

const (
	LogErr   = "error"
	LogWarn  = "warning"
	LogInfo  = "info"
	LogDebug = "debug"
	LogTest  = "test"
)

// Standard logger interface.
// Every logger should implement this interface.
// To be used by the framework.
type Logger interface {
	// Set the standard output
	SetStd(std *os.File)
	// Set the standard output for errors
	SetStdErr(std *os.File)
	// Write a message
	Write(t string, message string, args ...any)
	// Write an error message
	Error(msg any, args ...any)
	// Write a warning message
	Warning(msg string, args ...any)
	// Write an info message
	Info(msg string, args ...any)
	// Write a debug message
	Debug(msg string, args ...any)
	// Write a test message
	Test(msg string, args ...any)
}
