package utils

import (
    "os"
    "strings"
)

// Color codes for CLI output
const (
    Reset   = "\033[0m"
    Red     = "\033[31m"
    Green   = "\033[32m"
    Yellow  = "\033[33m"
    Blue    = "\033[34m"
    Cyan    = "\033[36m"
    White   = "\033[97m"
    Bold    = "\033[1m"
    Magenta = "\033[35m"
)

// Custom logger structure
type CustomLogger struct {
    prefix string
    color  string
}

// NewLogger creates a new custom logger
func NewLogger(prefix, color string) *CustomLogger {
    return &CustomLogger{
        prefix: prefix,
        color:  color,
    }
}

// WriteString writes directly to stdout with color
func (l *CustomLogger) WriteString(msg string) {
    output := []byte(l.color + l.prefix + msg + Reset + "\n")
    os.Stdout.Write(output)
}

// PrintBanner - Fancy CLI welcome screen
func PrintBanner() {
    banner := []string{
        "",
        Cyan + "────────────────────────────────────────" + Reset,
        White + "     ◽◾ " + Bold + "Primix CLI" + Reset + " v1.0 ◾◽",
        Cyan + "────────────────────────────────────────" + Reset,
        "",
        Magenta + "🚀 The Go Web Framework of the Future" + Reset,
        "",
    }
    
    output := strings.Join(banner, "\n")
    os.Stdout.Write([]byte(output + "\n"))
}

// Logger instances
var (
    InfoLogger    = NewLogger("[INFO] ", Green)
    WarnLogger    = NewLogger("[WARN] ", Yellow)
    ErrorLogger   = NewLogger("[ERROR] ", Red)
    SuccessLogger = NewLogger("[SUCCESS] ", Yellow)
)

// Logging functions
func Info(msg string) {
    InfoLogger.WriteString(msg)
}

func Warn(msg string) {
    WarnLogger.WriteString(msg)
}

func Error(msg string) {
    ErrorLogger.WriteString(msg)
}

func Success(msg string) {
    SuccessLogger.WriteString(msg)
}