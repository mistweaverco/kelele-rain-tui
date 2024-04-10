package main

import (
	"github.com/charmbracelet/log"
)

var VERSION = "1.0.0"

func main() {
	log.Info("Starting kelele rain 🌧️", "version", VERSION)
	overview()
}