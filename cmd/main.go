package main

import (
	"log"

	"github.com/perfectogo/live-chat/api"
	"github.com/perfectogo/live-chat/api/websocket"
)

func main() {
	go websocket.BroadCast()
	if err := api.New().Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
