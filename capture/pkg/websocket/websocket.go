package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Server struct {
	C           *websocket.Conn
	Recv        chan []byte
	Send        chan []byte
	MessageType int
}

func (s *Server) ServeWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Connected to: ", ws.RemoteAddr())
	s.C = ws
    s.MessageType = websocket.TextMessage
	// go func() {
	// 	for {
	// 		_, msg, err := ws.ReadMessage()
	// 		if err != nil {
	// 			fmt.Println("read message error")
	// 			return
	// 		}
	// 		fmt.Printf("Client: %s\n", msg)
	// 	}
	// }()
	//
	// go func() {
	// 	for msg := range s.Send {
	// 		err := s.C.WriteMessage(websocket.TextMessage, msg)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}
	// }()
}

func (s *Server) InitServer(port string) {
	http.HandleFunc("/", s.ServeWS)
	http.ListenAndServe(":"+port, nil)
}
