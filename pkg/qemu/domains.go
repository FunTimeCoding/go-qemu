package qemu

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-qemu/pkg/qemu/domain"
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

func (c *Client) Domains() []*domain.Domain {
	var result []*libvirtxml.Domain

	for _, element := range c.domains() {
		result = append(result, LoadDomain(&element))
	}

	return domain.NewSlice(result)
}

func (c *Client) domains() []libvirt.Domain {
	result, e := c.virt.ListAllDomains(0)
	errors.PanicOnError(e)

	return result
}
