package main

import (
	"fmt"
	"log"
	"net-sniffer/sniffer"
	"net/http"
)

func main() {
	fmt.Println("Starting network packet analyzer...")

	// start packet capturing
	go sniffer.StartSniffing()

	port := "8080"
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
