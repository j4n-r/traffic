package capture

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/j4n-r/traffic/pkg/websocket"
)

type Packet struct {
	version   byte
	protocol  byte
	timestamp uint32
	addr_type byte
	src_addr  net.IP
	src_port  uint16
	dst_addr  net.IP
	dst_port  uint16
}

func (p *Packet) String() string {
	return fmt.Sprintf(
		"Version: 0x%02X Protocol: 0x%02X Timestamp: %d Addr Type: 0x%02X Source Address: %s Source Port: %d Destination Address: %s Destination Port: %d\n",
		p.version,
		p.protocol,
		p.timestamp,
		p.addr_type,
		p.src_addr.String(),
		p.src_port,
		p.dst_addr.String(),
		p.dst_port,
	)
}

func (p *Packet) constructPayload() ([]byte, error) {
	var buf []byte
	// Append fields and check for errors
	var err error
	buf, err = binary.Append(buf, binary.BigEndian, p.version)
	if err != nil {
		return buf, err
	}

	buf, err = binary.Append(buf, binary.BigEndian, p.protocol)
	if err != nil {
		return buf, err
	}

	buf, err = binary.Append(buf, binary.BigEndian, p.timestamp)
	if err != nil {
		return buf, err
	}

	buf, err = binary.Append(buf, binary.BigEndian, p.addr_type)
	if err != nil {
		return buf, err
	}

	buf, err = binary.Append(buf, binary.BigEndian, p.src_addr)
	if err != nil {
		return buf, err
	}

	buf, err = binary.Append(buf, binary.BigEndian, p.src_port)
	if err != nil {
		return buf, err
	}

	buf, err = binary.Append(buf, binary.BigEndian, p.dst_addr)
	if err != nil {
		return buf, err
	}

	buf, err = binary.Append(buf, binary.BigEndian, p.dst_port)
	if err != nil {
		return buf, err
	}

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

	// add missing data
	p.version = 0x01

	if p.src_addr.To4() != nil {
		p.addr_type = 0x01
	} else {
		p.addr_type = 0x00
	}

	p.timestamp = uint32(time.Now().Unix())

	fmt.Println(p.String())
	switch layerType {
	case layers.LayerTypeDNS:
	default:
		if s.C != nil {
			buf, err := p.constructPayload()
			if err != nil {
				log.Fatal(1)
			}
			s.C.WriteMessage(s.MessageType, buf)
		}
	}
}
