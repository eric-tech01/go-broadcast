package lib

import (
	"net"
	"regexp"
)

var (
	ipReg   *regexp.Regexp
	addrReg *regexp.Regexp
)

func init() {
	ipReg, _ = regexp.Compile(`((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}`)
	addrReg, _ = regexp.Compile(`((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}:(([2-9]\d{3})|([1-5]\d{4})|(6[0-4]\d{3})|(65[0-4]\d{2})|(655[0-2]\d)|(6553[0-5]))`)

}

func GetNetInterface(ip string) (*net.Interface, error) {
	if ip == "127.0.0.1" {
		return nil, nil
	}
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				ipv4 := ipReg.FindString(address.String())
				if ipv4 == ip {
					ifi := &netInterfaces[i]
					return ifi, nil
				}
			}
		}
	}
	return nil, nil
}
