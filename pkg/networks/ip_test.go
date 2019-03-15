package networks

import (
	"net"
	"testing"
)

func TestIPv4ToInt(t *testing.T) {
	for _, c := range []struct {
		in   net.IP
		want uint32
	}{
		{net.ParseIP("192.168.1.1"), 3232235777},
		{net.ParseIP("0.0.0.0"), 0},
		{net.ParseIP("8.8.8.8"), 134744072},
		{net.ParseIP("255.255.255.255"), 4294967295},
	} {
		got := IPv4ToInt(c.in)
		if got != c.want {
			t.Errorf("IPv4ToInt(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIntToIPv4(t *testing.T) {
	for _, c := range []struct {
		in   uint32
		want net.IP
	}{
		{3232235777, net.ParseIP("192.168.1.1")},
		{0, net.ParseIP("0.0.0.0")},
		{134744072, net.ParseIP("8.8.8.8")},
		{4294967295, net.ParseIP("255.255.255.255")},
	} {
		got := IntToIPv4(c.in)
		if !got.Equal(c.want) {
			t.Errorf("IntToIPv4(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
