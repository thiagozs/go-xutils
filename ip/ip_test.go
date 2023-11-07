package ip

import (
	"testing"
)

func TestIPv4(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"192.168.1.1", true},
		{"255.255.255.255", true},
		{"0.0.0.0", true},
		{"256.0.0.0", false},
		{"192.168.1", false},
		{"192.168.1.256", false},
		{"abc.def.ghi.jkl", false},
		{"", false},
	}

	ipChecker := New()
	for _, test := range tests {
		if res := ipChecker.IPv4(test.ip); res != test.expected {
			t.Errorf("IPv4(%s) = %v; want %v", test.ip, res, test.expected)
		}
	}
}

func TestIPv6(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"2001:db8:85a3:0:0:8a2e:370:7334", true},
		{"2001:db8:85a3::8a2e:370:7334", true},
		{"::1", true},
		{"::", true},
		{"2001:db8::8a2e:37023:7334", false},
		{"2001::25de::cade", false},
		{"192.168.0.1", false},
		{"", false},
	}

	ipChecker := New()
	for _, test := range tests {
		if res := ipChecker.IPv6(test.ip); res != test.expected {
			t.Errorf("IPv6(%s) = %v; want %v", test.ip, res, test.expected)
		}
	}
}

func TestIP(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"127.0.0.1", true},
		{"::1", true},
		{"2001:db8::8a2e", true},
		{"256.0.0.1", false},
		{"192.168.0", false},
		{"2001::25de::cade", false},
		{"", false},
	}

	ipChecker := New()
	for _, test := range tests {
		if res := ipChecker.IP(test.ip); res != test.expected {
			t.Errorf("IP(%s) = %v; want %v", test.ip, res, test.expected)
		}
	}
}
