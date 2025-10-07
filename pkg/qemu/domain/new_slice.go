package domain

import "libvirt.org/go/libvirtxml"

func NewSlice(input []*libvirtxml.Domain) []*Domain {
	var result []*Domain

	for _, element := range input {
		result = append(result, New(element))
	}

	return result
}
