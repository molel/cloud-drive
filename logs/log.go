package logs

import (
	"log"
	"os"
	"strings"
)

func init() {
	file, err := os.OpenFile("./logs/logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("cannot open logs file:\n%s", err.Error())
	}
	log.SetOutput(file)
	log.SetPrefix(strings.Repeat("-", 80) + "\n")
	log.Println("Session start")
	log.SetPrefix("\n")
}
