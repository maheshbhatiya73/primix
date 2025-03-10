package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func Init() {
	Logger = log.New(os.Stdout, "PRIMIX: ", log.LstdFlags)
}
