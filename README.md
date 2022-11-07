# simplelog!

## Installation
Easily install github.com/Nigel2392/simplelog with the following command:
```
go get github.com/Nigel2392/simplelog
```

## Usage

Simple import the package, and initiate the logger with the following code:
```go
import "github.com/Nigel2392/simplelog"
var logger = simplelog.NewLogger(loglevel)
// Set on error std
logger.SetStd(os.File)
logger.SetStdErr(os.File) // For errors, optional
```

## Log levels
The logger uses the following log levels:
It gets the integer value of the log level, and logs everything above that level.
```go
error   = 4 = "error" 
warning = 3 = "warning" 
info    = 2 = "info" 
debug   = 1 = "debug" 
test    = 0 = "test" 
```

## Logger interface
The logger interface is as follows:
```go
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
```