package networks

import (
	"net"
	"sort"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/pkg/errors"
)

func New(cidrblock string) (*Network, error) {
	_, ipnet, err := net.ParseCIDR(cidrblock)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse cidr")
	}
	first, last := cidr.AddressRange(ipnet)
	return &Network{
		Cidr:     ipnet,
		First:    &first,
		Last:     &last,
		FirstInt: IPv4ToInt(first),
		LastInt:  IPv4ToInt(last),
	}, nil
}

type Network struct {
	Cidr     *net.IPNet
	First    *net.IP
	Last     *net.IP
	FirstInt uint32
	LastInt  uint32
}

func Sort(networks []*Network) []*Network {
	sort.Slice(networks, func(i, j int) bool {
		return networks[i].FirstInt < networks[j].FirstInt
	})
	return networks
}
