package capture

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/j4n-r/traffic/pkg/websocket"
)

type Packet struct {
	src_addr string
	src_port string
	dst_addr string
	dst_port string
}

func StartCapture(device string, s *websocket.Server) {
	// Open a live capture handle on the "tailscale0" interface
	handle, err := pcap.OpenLive(device, 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set up a packet source to read packets from the handle
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		handlePacket(packet.Data(), s) // Process each packet as a byte slice
	}
}

func handlePacket(packetData []byte, s *websocket.Server) {
	p := Packet{}
	packet := gopacket.NewPacket(packetData, layers.LayerTypeIPv4, gopacket.Default)

	if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		p.src_addr = ip.SrcIP.String()
		p.dst_addr = ip.DstIP.String()
	} else if ipLayer := packet.Layer(layers.LayerTypeIPv6); ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv6)
		p.src_addr = ip.SrcIP.String()
		p.dst_addr = ip.DstIP.String()
	}

	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		// Get actual TCP data from this layer and print source/destination ports
		tcp, _ := tcpLayer.(*layers.TCP)
		p.src_port = tcp.SrcPort.String()
		p.dst_port = tcp.DstPort.String()
	}

	var layerType gopacket.LayerType
	for _, layer := range packet.Layers() {
		layerType = layer.LayerType()
	}
	switch layerType {
	case layers.LayerTypeDNS:
	default:
		fmt.Printf("%s src: %s:%s dst:%s:%s\n", layerType, p.src_addr, p.src_port, p.dst_addr, p.dst_port)
		s.C.WriteMessage(s.MessageType, []byte(fmt.Sprintf("%s src: %s:%s dst:%s:%s\n", layerType, p.src_addr, p.src_port, p.dst_addr, p.dst_port)))
	}
}
