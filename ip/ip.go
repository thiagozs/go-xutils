package ip

import (
	"net"
	"strconv"
	"strings"
)

type Ip struct{}

func New() *Ip {
	return &Ip{}
}

func (i *Ip) IPv4(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}
	for _, part := range parts {
		if len(part) == 0 || len(part) > 3 {
			return false
		}
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func (i *Ip) IPv6(ip string) bool {
	return net.ParseIP(ip) != nil && strings.Contains(ip, ":")
}

func (i *Ip) IP(ip string) bool {
	return i.IPv4(ip) || i.IPv6(ip)
}
