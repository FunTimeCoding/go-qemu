package domain

import "libvirt.org/go/libvirtxml"

func New(d *libvirtxml.Domain) *Domain {
	return &Domain{Name: d.Name, Raw: d}
}
