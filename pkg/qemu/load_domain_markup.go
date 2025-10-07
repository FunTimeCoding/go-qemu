package qemu

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

func LoadDomain(d *libvirt.Domain) *libvirtxml.Domain {
	result := &libvirtxml.Domain{}
	errors.PanicOnError(result.Unmarshal(DomainMarkup(d)))

	return result
}
