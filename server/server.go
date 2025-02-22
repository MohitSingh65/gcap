package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net-sniffer/sniffer"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type PacketData struct {
	Timestamp string `json:"timestamp"`
	SrcMAC    string `json:"src_mac"`
	DstMAC    string `json:"dst_mac"`
	SrcIP     string `json:"src_ip"`
	DstIP     string `json:"dst_ip"`
	SrcPort   uint16 `json:"src_port"`
	DstPort   uint16 `json:"dst_port"`
	Payload   string `json:"payload"`
}

func main() {
	http.HandleFunc("/start", startCaptureHandler)
	http.HandleFunc("/stop", stopCaptureHandler)
	http.HandleFunc("/ws", wsHandler)

	log.Println("Starting ")
}
