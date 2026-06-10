package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("WS request: host=%s origin=%s", r.Host, r.Header.Get("Origin"))

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{
			"localhost:*",
			"127.0.0.1:*",
		},
	})
	if err != nil {
		log.Println("websocket accept error:", err)
		return
	}
	defer conn.CloseNow()

	log.Println("client connected")

	ctx := context.Background()

	for {
		messageType, data, err := conn.Read(ctx)
		if err != nil {
			log.Println("read error / client disconnected:", err)
			return
		}

		log.Println("received:", string(data))

		writeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		err = conn.Write(writeCtx, messageType, data)
		cancel()

		if err != nil {
			log.Println("write error:", err)
			return
		}
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("/ws", wsHandler)

	log.Println("Go server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}