package main

import (
	"github.com/j4n-r/traffic/pkg/capture"   // Import the capture package
	"github.com/j4n-r/traffic/pkg/websocket" // Import the websocket package
)

func main() {
	s := websocket.Server{
		Recv: make(chan []byte),
		Send: make(chan []byte),
	}
	go s.InitServer("8999")
	capture.StartCapture("tailscale0", &s)
}
