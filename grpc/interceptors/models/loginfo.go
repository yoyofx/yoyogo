package models

type LoggerInfo struct {
	StartTime string
	Status    int
	Duration  string
	HostName  string
	Method    string
	Path      string
	Request   interface{}
	Response  interface{}
}

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func (p *LoggerInfo) StatusCodeColor() string {
	switch p.Status {
	case 200:
		return green
	case 500:
		return red
	default:
		return yellow
	}
}

// MethodColor is the ANSI color for appropriately logging http method to a terminal.
func (p *LoggerInfo) MethodColor() string {
	switch p.Method {
	case "unary":
		return blue
	case "stream":
		return green
	default:
		return reset
	}
}

// ResetColor resets all escape attributes.
func (p *LoggerInfo) ResetColor() string {
	return reset
}
