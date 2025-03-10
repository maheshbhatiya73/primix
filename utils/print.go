package utils

import (
	"fmt"
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

// PrintBanner - Fancy CLI welcome screen
func PrintBanner() {
	fmt.Println("")
	fmt.Println(Cyan + "────────────────────────────────────────" + Reset)
	fmt.Println(White + "     ◽◾ " + Bold + "Primix CLI" + Reset + " v1.0 ◾◽")
	fmt.Println(Cyan + "────────────────────────────────────────" + Reset)
	fmt.Println("")
	fmt.Println(Magenta + "🚀 The Go Web Framework of the Future" + Reset)
	fmt.Println("")
}

// Info - Green info message
func Info(msg string) {
	fmt.Println(Green + "[INFO] " + msg + Reset)
}

// Warn - Yellow warning message
func Warn(msg string) {
	fmt.Println(Yellow + "[WARN] " + msg + Reset)
}

// Error - Red error message
func Error(msg string) {
	fmt.Println(Red + "[ERROR] " + msg + Reset)
}

// Success - Yellow success message
func Success(msg string) {
	fmt.Println(Yellow + "[SUCCESS] " + msg + Reset)
}
