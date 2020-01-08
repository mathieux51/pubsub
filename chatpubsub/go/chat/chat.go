package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func run() error {
	http.Handle("/", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
