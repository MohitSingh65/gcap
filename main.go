package main

import (
	"fmt"
	"net-sniffer/server"
	"net-sniffer/sniffer"
)

func main() {
	fmt.Println("Starting network packet analyzer...")

	// Start websocket server in a separate Goroutine
	go server.StartWebSocketServer()

	// start packet capturing
	go sniffer.StartSniffing()

	select {}
}
