package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var localIP = "192.168.10.102"
var size = 0

func main() {
	//  获取 libpcap 的版本
	version := pcap.Version()
	fmt.Println(version)
	//  获取网卡列表
	var device string
	device = findNetName(localIP)
	fmt.Println(device)
	hand, err := pcap.OpenLive(device, int32(999999999), true, -1*time.Second)
	fmt.Println(err)
	hand.SetBPFFilter("port 8080")

	defer hand.Close()
	// mac,err:=findMacAddrByIp("192.168.10.102")
	packetSource := gopacket.NewPacketSource(hand, hand.LinkType())
	// packet,_:=packetSource.NextPacket()
	// for packet:=range packetSource.Packets(){
	//	ipLayer := packet.TransportLayer()
	//	if ipLayer!=nil{
	//		ipLayer.LayerPayload()
	//	}
	//	tcp,_:=ipLayer.(*layers.TCP)
	//	fmt.Println(len(tcp.Payload))
	//	fmt.Println(len(tcp.Contents))
	//	fmt.Println(tcp.DataOffset)
	//	fmt.Println(packet.String())
	//	//packet.Data()
	//	//packet.
	// }

	for packet := range packetSource.Packets() {
		//  解析 IP 层
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			//  解析 TCP 层
			tcpLayer := packet.Layer(layers.LayerTypeTCP)
			if tcpLayer != nil {
				tcp, _ := tcpLayer.(*layers.TCP)
				if len(tcp.Payload) > 0 {
					ip, _ := ipLayer.(*layers.IPv4)

					if ip.DstIP.String() == localIP && tcp.DstPort.String() == "8080" {
						size += len(tcp.Payload)
						fmt.Println(size)
					}
					fmt.Printf("%s:%s->%s:%s\n%s\n",
						ip.SrcIP, tcp.SrcPort,
						ip.DstIP, tcp.DstPort,
						string(tcp.Payload))
				}
			} else if errLayer := packet.ErrorLayer(); errLayer != nil {
				fmt.Printf("tcp.err: %v", errLayer)
			}
		} else if errLayer := packet.ErrorLayer(); errLayer != nil {
			fmt.Printf("ip.err: %v", errLayer)
		}
	}
}

func findNetName(prefix string) string {
	//  获取网卡列表
	var devices []pcap.Interface
	devices, _ = pcap.FindAllDevs()
	for _, d := range devices {
		for _, addr := range d.Addresses {
			if ip4 := addr.IP.To4(); ip4 != nil {
				if strings.HasPrefix(ip4.String(), prefix) {
					data, _ := json.MarshalIndent(d, "", "  ")
					fmt.Println(string(data))
					return d.Name
				}
			}
		}
	}
	return ""
}

// 获取网卡的IPv4地址
func findDeviceIpv4(device pcap.Interface) string {
	for _, addr := range device.Addresses {
		if ipv4 := addr.IP.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	panic("device has no IPv4")
}

// 根据网卡的IPv4地址获取MAC地址
// 有此方法是因为gopacket内部未封装获取MAC地址的方法，所以这里通过找到IPv4地址相同的网卡来寻找MAC地址
func findMacAddrByIp(ip string) (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(interfaces)
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}

		for _, addr := range addrs {
			if a, ok := addr.(*net.IPNet); ok {
				if ip == a.IP.String() {
					return i.HardwareAddr.String(), nil
				}
			}
		}
	}
	return "", errors.New(fmt.Sprintf("no device has given ip: %s", ip))
}
