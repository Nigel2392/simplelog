package simplelog

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

// Preset colors for use in the logger's Colorize function
var (

	// LogTest
	ColorLevelTest = Purple

	// LogDebug
	ColorLevelDebug = Green

	// LogInfo
	ColorLevelInfo = Blue

	// LogWarn
	ColorLevelWarning = Yellow

	// LogErr
	ColorLevelError = Red

	// No level, default switch case opt.
	ColorNoLevel = Green
)

// Colorize a message based on the loglevel
func Colorize(level int, msg string) string {
	var selected string
	switch level {
	case 0:
		selected = ColorLevelTest
	case 1:
		selected = ColorLevelDebug
	case 2:
		selected = ColorLevelInfo
	case 3:
		selected = ColorLevelWarning
	case 4:
		selected = ColorLevelError
	default:
		selected = ColorNoLevel
	}
	return selected + msg + Reset
}
