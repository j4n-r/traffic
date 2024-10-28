package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Open a live capture handle on the "tailscale0" interface
	handle, err := pcap.OpenLive("tailscale0", 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set up a packet source to read packets from the handle
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		handlePacket(packet.Data()) // Process each packet as a byte slice
	}
}

func handlePacket(packetData []byte) {
	packet := gopacket.NewPacket(packetData, layers.LayerTypeIPv4, gopacket.Default)

	// Get the TCP layer from this packet (if it exists)
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		fmt.Println("This is a TCP packet!")

		// Get actual TCP data from this layer and print source/destination ports
		tcp, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("From src port %d to dst port %d\n", tcp.SrcPort, tcp.DstPort)
	}

	// Iterate over all layers and print out each layer type
	for _, layer := range packet.Layers() {
		fmt.Println("PACKET LAYER:", layer.LayerType())
	}
}
