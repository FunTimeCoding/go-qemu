package qemu

import (
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-qemu/pkg/qemu/hardware_address"
	"net"
	"sort"
)

func (c *Client) HardwareAddresses() []net.HardwareAddr {
	var result []net.HardwareAddr

	for _, element := range c.Domains() {
		for _, inner := range element.Raw.Devices.Interfaces {
			result = append(
				result,
				network.PhysicalAddress(inner.MAC.Address),
			)
		}
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return hardware_address.LessThan(result[i], result[j])
		},
	)

	return result
}
