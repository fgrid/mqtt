package main

import (
	"log"

	"github.com/fgrid/mqtt"
)

func main() {
	c := mqtt.DefaultClient
	if err := c.Connect("tcp", "iot.eclipse.org:1883"); err != nil {
		log.Fatalf("could not connect client: %s", err.Error())
	} else {
		log.Printf("client connected to server")
	}
}
