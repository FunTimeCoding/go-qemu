package qemu

import (
	"github.com/digitalocean/go-qemu/qmp"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
)

func (c *Client) Monitor(domain string) *qmp.LibvirtRPCMonitor {
	q, e := net.DialTimeout("unix", c.socket, c.timeout)
	errors.PanicOnError(e)

	result := qmp.NewLibvirtRPCMonitor(domain, q)
	errors.PanicOnError(result.Connect())

	c.monitors = append(c.monitors, result)

	return result
}
