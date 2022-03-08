package logger

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Info  = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	Debug = log.New(ioutil.Discard, "DBUG: ", log.LstdFlags)
	Error = log.New(os.Stderr, "ERRO: ", log.LstdFlags)
)

func UpdateENV() {
	if os.Getenv("SCRAPER_LOG") == "DEBUG" {
		Debug.SetOutput(os.Stdout)
	}
}

func init() {
	UpdateENV()
}
