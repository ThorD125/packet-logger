package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"net/http"
	"sync"
)

var (
	clients     []chan gopacket.Packet
	clientMu    sync.Mutex
	isPaused    bool
	pauseResume sync.Mutex
	handle      *pcap.Handle

	ipv4AddrOfInterface string
)

func main() {

	http.HandleFunc("/", handleSummonWebpage)
	http.HandleFunc("/updatePackets", handleUpdatePackets)
	http.HandleFunc("/pause", handlePause)
	http.HandleFunc("/resume", handleResume)
	http.HandleFunc("/selectDevice", handleSelectDevice)

	http.ListenAndServe(":8888", nil)

}

type Payload struct {
	Field1 uint8
	Field2 uint8
	Field3 uint8
	Field4 uint8
}
