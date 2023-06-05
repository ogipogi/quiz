package log

import "log"

func Fatal(message ...string) {
	log.Fatal("FATAL:", message)
}

func Error(message ...string) {
	log.Println("ERROR:", message)
}

func Warn(message ...string) {
	log.Println("WARNING:", message)
}

func Info(message ...string) {
	log.Println("WARNING:", message)
}
