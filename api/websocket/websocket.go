package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Content       string `json:"content"`
	UserName      string `json:"username"`
	GeterUserName string `json:"geterUsername"`
}

var (
	Clients  = make(map[*websocket.Conn]bool)
	Messages = make(chan Message)
)

func WebSocket(ctx *gin.Context) {
	//
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	defer conn.Close()
	if err != nil {
		log.Println("-1-")
		log.Fatal(err)
	}
	Clients[conn] = true

	//
	for {
		message := Message{}

		err := conn.ReadJSON(&message)

		if err != nil {
			log.Println(err, "-2-")
			//	log.Fatal(err)
			delete(Clients, conn)
			break
		}

		Messages <- message
	}
}
