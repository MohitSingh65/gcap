package sniffer

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"net-sniffer/server"
)

// Start capturing the packets
func StartSniffing() {
	device := "wlan0"
	snapshotlen := int32(65535)
	promiscuous := false
	timeout := time.Second * 30
	handle, err := pcap.OpenLive(device, snapshotlen, promiscuous, timeout)
	if err != nil {
		log.Fatalf("Error opening device %s: %v", device, err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Read packets in loop
	for packet := range packetSource.Packets() {
		packetData := fmt.Sprintf("Captured packet: %v", packet)
		fmt.Println(packetData)
		server.BroadcastMessage(packetData) // Send to websocket client
	}
}
