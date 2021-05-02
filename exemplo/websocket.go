package exemplo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1204,
	WriteBufferSize: 1204,
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "welcome")
	})
	http.HandleFunc("/ws", func(resp http.ResponseWriter, req *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(resp, req, nil)
		if err != nil {
			log.Println(err)
		}
		log.Println("client connected")
		reader(ws)
	})
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		log.Println("message from client: " + string(p))

		msg := []byte("Let's start to talk something.")
		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

// func main() {
// 	setupRoutes()
// 	http.ListenAndServe(":8080", nil)
// }
