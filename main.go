package logger

import "fmt"

var (
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[34m"
	magenta = "\x1b[35m"
	cyan    = "\x1b[36m"
	white   = "\x1b[37m"
	reset   = "\x1b[0m"
)

func Warn(msg string) {
	fmt.Println(yellow + "WARNING: " + reset + msg)
}
func Info(msg string) {
	fmt.Println(green + "INFO: " + reset + msg)
}
func Error(msg string) {
	fmt.Println(red + "ERROR: " + reset + msg)
}
