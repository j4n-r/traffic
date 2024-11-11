package capture

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/j4n-r/traffic/pkg/websocket"
)

type Packet struct {
	version   byte
	protocol  layers.ProtocolFamily // byte
	timestamp uint32
	src_addr  net.IP
	src_port  net.IP
	dst_addr  []byte
	dst_port  uint16
}

func constructPayload(buf []byte, p Packet) ([]byte, error) {
	p.version = 0x01
    
    if p.src_addr.To4() != nil {

    }
    
	binary.Append(buf, binary.BigEndian, p.version)
	return buf, nil
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
		p.src_addr = ip.SrcIP
		p.dst_addr = ip.DstIP
	} else if ipLayer := packet.Layer(layers.LayerTypeIPv6); ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv6)
		p.src_addr = ip.SrcIP
		p.dst_addr = ip.DstIP
	}

	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		// Get actual TCP data from this layer and print source/destination ports
		tcp, _ := tcpLayer.(*layers.TCP)
		p.src_port = uint16(tcp.SrcPort)
		p.dst_port = uint16(tcp.DstPort)
	}

	var layerType gopacket.LayerType
	for _, layer := range packet.Layers() {
		layerType = layer.LayerType()
	}
	switch layerType {
	case layers.LayerTypeDNS:
	default:
		fmt.Printf("%s src: %s:%s dst:%s:%s\n", layerType, p.src_addr, p.src_port, p.dst_addr, p.dst_port)
		if s.C != nil {
			s.C.WriteMessage(s.MessageType, []byte(fmt.Sprintf("%s src: %s:%s dst:%s:%s\n", layerType, p.src_addr, p.src_port, p.dst_addr, p.dst_port)))
		}
	}
}
