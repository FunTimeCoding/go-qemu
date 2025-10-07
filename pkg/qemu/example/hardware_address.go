package example

import (
	"fmt"
	"github.com/funtimecoding/go-qemu/pkg/qemu"
)

func HardwareAddress() {
	c := qemu.New()
	defer c.Close()

	for _, a := range c.HardwareAddresses() {
		fmt.Printf("Hardware address: %+v\n", a)
	}

	fmt.Printf("Next address: %+v\n", c.FindHardwareAddress())
}
