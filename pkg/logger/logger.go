package logger


import (
	"log"
	"strings"
)

var level string


func InitLogger(logLevel string) {
	level = strings.ToUpper(logLevel)
	log.Printf("Logger initialized with level: %s", level)
}

func Info(msg string) {
	if level == "INFO" || level == "DEBUG" {
		log.Printf("[INFO] %s", msg)
	}
}

func Debug(msg string) {
	if level == "DEBUG" {
		log.Printf("[DEBUG] %s", msg)
	}
}

func Error(msg string) {
	log.Printf("[ERROR] %s", msg)
}