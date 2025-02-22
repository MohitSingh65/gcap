package sniffer

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Start capturing the packets
func StartSniffing() {
	handle, err := pcap.OpenLive("wlan0", 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal("Error opening device:", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Read packets in loop
	for packet := range packetSource.Packets() {
		fmt.Println("Captured packet:", packet)
	}
}
