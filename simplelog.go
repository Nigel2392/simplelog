package simplelog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

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

func (l *logger) Write(t string, message string, args ...any) {
	var msg string
	if len(args) > 0 {
		msg = fmt.Sprintf(message, args...)
	} else {
		msg = message
	}
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
	case LogErr:
		return 4
	case LogWarn:
		return 3
	case LogInfo:
		return 2
	case LogDebug:
		return 1
	case LogTest:
		return 0
	default:
		return 1
	}
}

// Write a Error message to the logger.
func (l *logger) Error(msg any, args ...any) {
	var message string
	switch msg := msg.(type) {
	case string:
		message = msg
	case error:
		message = msg.Error()
	default:
		message = fmt.Sprint(msg)
	}
	l.Write(LogErr, message, args...)
}

// Write a Warning message to the logger.
func (l *logger) Warning(msg string, args ...any) {
	l.Write(LogWarn, msg, args...)
}

// Write a Info message to the logger.
func (l *logger) Info(msg string, args ...any) {
	l.Write(LogInfo, msg, args...)
}

// Write a Debug message to the logger.
func (l *logger) Debug(msg string, args ...any) {
	l.Write(LogDebug, msg, args...)
}

// Write a Test message to the logger.
func (l *logger) Test(msg string, args ...any) {
	l.Write(LogTest, msg, args...)
}

// Wrap a message with time
func WrapTime(t string, msg string) string {
	var time string = time.Now().Format("2006-01-02 15:04:05")
	var typ string = strings.ToUpper(t)
	return "[" + time + " " + typ + "] " + msg
}
