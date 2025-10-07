package qemu

import (
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-qemu/pkg/qemu/hardware_address"
	"net"
)

func (c *Client) FindHardwareAddress() net.HardwareAddr {
	return hardware_address.FirstUnused(
		network.PhysicalAddress(c.hardwareAddress),
		c.HardwareAddresses(),
	)
}
