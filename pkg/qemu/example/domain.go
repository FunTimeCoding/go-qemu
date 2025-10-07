package example

import (
	"fmt"
	"github.com/funtimecoding/go-qemu/pkg/qemu"
)

func Domain() {
	c := qemu.New()
	defer c.Close()

	for _, element := range c.Domains() {
		fmt.Printf("Domain: %+v\n", element.Raw.Name)

		for _, inner := range element.Raw.Devices.Interfaces {
			fmt.Printf("Hardware address: %+v\n", inner.MAC.Address)
		}
	}
}
