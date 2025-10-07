package domain

import "libvirt.org/go/libvirtxml"

func NewSlice(v []*libvirtxml.Domain) []*Domain {
	var result []*Domain

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
