package main

import (
	"compare_folders/handlers"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

// initLogger sets up the logger with a file output and a proper format
func initLogger() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to log to file: ", err)
	}
	log.SetOutput(file)
	log.SetLevel(logrus.InfoLevel)
}

func main() {
	initLogger()
	handlers.StartComparison(log) // Starts the folder comparison process
}
