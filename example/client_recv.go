package main

import (
	"fmt"
	"net"

	broadcast "github.com/eric-tech01/go-broadcast"
	"github.com/eric-tech01/go-broadcast/lib"
	"golang.org/x/net/ipv4"
)

func main() {
	ip := "192.168.0.48"
	broadcastIp := "239.0.0.1:12345"
	ifn, err := lib.GetNetInterface(ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v \n", ifn)

	err = broadcast.Receive(broadcastIp, ip, ifn, msgHandler)
	fmt.Println(err)
}

func msgHandler(cm *ipv4.ControlMessage, src net.Addr, n int, b []byte) {
	fmt.Printf("recv len: %d \n", n)
	fmt.Printf("%s \n", b[0:n])

}
