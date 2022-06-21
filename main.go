package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yshurik/echopoint/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
