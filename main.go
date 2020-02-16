package main

import (
	"ayls/blockchain-dns-server/dns"
	"fmt"
	"github.com/natefinch/lumberjack"
	"log"
	"os"
)

var logger *log.Logger

func main() {
	e, err := os.OpenFile("./log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	logger = log.New(e, "", log.Ldate|log.Ltime)
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "./log.txt",
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
	})

	dns.Start(logger)
}
