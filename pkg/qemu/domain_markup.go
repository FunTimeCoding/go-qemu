package qemu

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"libvirt.org/go/libvirt"
)

func DomainMarkup(d *libvirt.Domain) string {
	// This would cause the log message in libvirtd if Free() is not called:
	// 2023-07-28 09:41:22.218+0000: 8364973568: error : virNetSocketReadWire:1785 : End of file while reading data: Input/output error
	result, e := d.GetXMLDesc(0)
	errors.PanicOnError(e)
	errors.PanicOnError(d.Free())

	return result
}
