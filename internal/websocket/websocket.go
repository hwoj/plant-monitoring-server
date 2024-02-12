package websocket

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	clients = make(map[*websocket.Conn]bool)
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	defer func() {
		conn.Close()
		delete(clients, conn)
	}()

	clients[conn] = true
	TestUpdate()
}

func TestUpdate() {
	for {
		dataPoint := time.Now().Format("15:04:05")

		BroadcastUpdate(dataPoint)
		time.Sleep(time.Second)
	}
}

func BroadcastUpdate(dataPoint string) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(dataPoint))

		if err != nil {
			log.Println(err)
			client.Close()
			delete(clients, client)
		}
	}
}
