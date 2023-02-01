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
const (
	// LogTest
	DefaultLevel0 = Purple
	// LogDebug
	DefaultLevel1 = Green
	// LogInfo
	DefaultLevel2 = Blue
	// LogWarn
	DefaultLevel3 = Yellow
	// LogErr
	DefaultLevel4 = Red
	// No level, default switch case opt.
	DefaultNoLevel = Green
)

// Colorize a message based on the loglevel
func Colorize(level int, msg string) string {
	var selected string
	switch level {
	case 0:
		selected = DefaultLevel0
	case 1:
		selected = DefaultLevel1
	case 2:
		selected = DefaultLevel2
	case 3:
		selected = DefaultLevel3
	case 4:
		selected = DefaultLevel4
	default:
		selected = DefaultNoLevel
	}
	return selected + msg + Reset
}
