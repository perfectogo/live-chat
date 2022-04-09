package websocket

import (
	"log"
)

func BroadCast() {
	for {
		message := <-Messages
		log.Println(message)

		for client := range Clients {
			err := client.WriteJSON(message)
			if err != nil {
				delete(Clients, client)
				break
			}
		}
	}
}
