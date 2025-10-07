package example

import "github.com/funtimecoding/go-qemu/pkg/qemu"

func DeleteDomain() {
	c := qemu.New()
	defer c.Close()

	c.DeleteDomain(c.Domain())
}
