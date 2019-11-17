package main

import(
	"os"
	"log"
)

var (
	logFile, _ = os.OpenFile("url_shortener.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	logger =  log.New(logFile, "", log.Ldate | log.Ltime)
)