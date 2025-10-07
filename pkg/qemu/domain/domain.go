package domain

import "libvirt.org/go/libvirtxml"

type Domain struct {
	Name string
	Raw  *libvirtxml.Domain
}
