package qemu

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"libvirt.org/go/libvirt"
)

func IsPersistent(d *libvirt.Domain) bool {
	result, e := d.IsPersistent()
	errors.PanicOnError(e)

	return result
}
