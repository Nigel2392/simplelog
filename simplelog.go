package simplelog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ANSI color codes
const (
	Reset        string = "\033[0m"
	Red          string = "\033[31m"
	Green        string = "\033[32m"
	Yellow       string = "\033[33m"
	Blue         string = "\033[34m"
	Purple       string = "\033[35m"
	Cyan         string = "\033[36m"
	White        string = "\033[37m"
	BrightRed    string = "\033[31;1m"
	BrightGreen  string = "\033[32;1m"
	BrightYellow string = "\033[33;1m"
	BrightBlue   string = "\033[34;1m"
	BrightPurple string = "\033[35;1m"
	BrightCyan   string = "\033[36;1m"
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
	Write(t string, message ...any)
	// Write an error message
	Error(msg ...any)
	// Write a warning message
	Warning(msg ...any)
	// Write an info message
	Info(msg ...any)
	// Write a debug message
	Debug(msg ...any)
	// Write a test message
	Test(msg ...any)
}

// Default logger
type logger struct {
	Level  string `json:"level"`
	stdout *os.File
	stderr *os.File
}

func NewLogger(level string) *logger {
	return &logger{
		Level:  level,
		stdout: os.Stdout,
		stderr: os.Stdout,
	}
}

// Set standard output
func (l *logger) SetStd(std *os.File) {
	l.stdout = std
}

// Set standard output for errors
func (l *logger) SetStdErr(std *os.File) {
	l.stderr = std
}

// Generate a message
func (l *logger) getMessage(t string, msg string) string {
	if l.GetLevelFromType(t) >= 4 {
		if l.stderr == os.Stdout {
			return Colorize(l.GetLevelFromType(t), WrapTime(t, msg))
		} else {
			return WrapTime(t, msg)
		}
	}
	if l.stdout == os.Stdout {
		if l.GetLevel() <= l.GetLevelFromType(t) {
			return Colorize(l.GetLevelFromType(t), WrapTime(t, msg))
		}
	} else {
		return WrapTime(t, msg)
	}
	return ""
}

func (l *logger) Write(t string, message ...any) {
	var msg string = fmt.Sprint(message...)
	var console_msg = l.getMessage(t, msg)
	if console_msg != "" {
		if l.GetLevelFromType(t) >= 4 {
			fmt.Fprintln(l.stderr, console_msg)
			return
		}
		fmt.Fprintln(l.stdout, console_msg)
	}
}

func (l *logger) GetLevel() int {
	return l.GetLevelFromType(l.Level)
}
func (l *logger) GetLevelFromType(t string) int {
	switch t {
	case "error":
		return 4
	case "warning":
		return 3
	case "info":
		return 2
	case "debug":
		return 1
	case "test":
		return 0
	default:
		return 1
	}
}

// Write a Error message to the logger.
func (l *logger) Error(msg ...any) {
	l.Write("error", msg...)
}

// Write a Warning message to the logger.
func (l *logger) Warning(msg ...any) {
	l.Write("warning", msg...)
}

// Write a Info message to the logger.
func (l *logger) Info(msg ...any) {
	l.Write("info", msg...)
}

// Write a Debug message to the logger.
func (l *logger) Debug(msg ...any) {
	l.Write("debug", msg...)
}

// Write a Test message to the logger.
func (l *logger) Test(msg ...any) {
	l.Write("test", msg...)
}

// Colorize a message based on the loglevel
func Colorize(level int, msg string) string {
	var selected string
	switch level {
	case 0:
		selected = Purple
	case 1:
		selected = Green
	case 2:
		selected = Blue
	case 3:
		selected = Yellow
	case 4:
		selected = Red
	default:
		selected = Green
	}
	return selected + msg + Reset
}

// Wrap a message with time
func WrapTime(t string, msg string) string {
	var time string = time.Now().Format("2006-01-02 15:04:05")
	var typ string = strings.ToUpper(t)
	return "[" + time + " " + typ + "] " + msg
}
