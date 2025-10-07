package qemu

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"libvirt.org/go/libvirt"
)

func Name(d *libvirt.Domain) string {
	result, e := d.GetName()
	errors.PanicOnError(e)

	return result
}
