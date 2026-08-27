package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils/v2/ip"
)

func main() {
	fmt.Println("IPv4:", ip.IsIPv4("192.0.2.10"))
	fmt.Println("IPv6:", ip.IsIPv6("2001:db8::1"))
	fmt.Println("IP válido:", ip.IsValid("2001:db8::1"))
}
