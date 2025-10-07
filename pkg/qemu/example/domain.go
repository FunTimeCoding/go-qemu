package example

import (
	"fmt"
	"github.com/funtimecoding/go-qemu/pkg/qemu"
)

func Domain() {
	c := qemu.New()
	defer c.Close()

	for _, d := range c.Domains() {
		fmt.Printf("Domain: %+v\n", d.Raw.Name)

		for _, inner := range d.Raw.Devices.Interfaces {
			fmt.Printf("Hardware address: %+v\n", inner.MAC.Address)
		}
	}
}
