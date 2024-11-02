package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
    "golang.org/x/net/websocket"
)

type ws struct{
    conn net.Conn
    buffwr net.Buffers
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8999", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
    
	if r.Header.Get("Upgrade") == "websocket" && r.Header.Get("Connection") == "Upgrade" {
		err := upgradeConn(w, r)

		if err != nil {
			return
		}

		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
        
	}

}

func upgradeConn(w http.ResponseWriter, r *http.Request) error {
	reqKey := r.Header.Get("Sec-WebSocket-Key")
	if len(reqKey) == 0 {
		return fmt.Errorf("No websockt key found")
	}

	resKey := computeResKey(reqKey)

	w.Header().Set("Upgrade", "websocket")
	w.Header().Set("Connection", "Upgrade")
	w.Header().Set("Sec-WebSocket-Accept", resKey)

	w.WriteHeader(http.StatusSwitchingProtocols)
	log.Printf("WebSocket connection successfully upgraded for client: %s", r.RemoteAddr)
	return nil

}

func computeResKey(reqKey string) (resKey string) {
	magicString := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	reqKey = reqKey + magicString
	h := sha1.New()

	_, err := io.WriteString(h, reqKey)
	if err != nil {
		return ""
	}
	sha1sum := h.Sum(nil)
	resKey = base64.StdEncoding.EncodeToString(sha1sum)
	return resKey
}
