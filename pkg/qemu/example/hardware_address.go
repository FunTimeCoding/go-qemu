package example

import (
	"fmt"
	"github.com/funtimecoding/go-qemu/pkg/qemu"
)

func HardwareAddress() {
	c := qemu.New()
	defer c.Close()

	for _, element := range c.HardwareAddresses() {
		fmt.Printf("Hardware address: %+v\n", element)
	}

	fmt.Printf("Next address: %+v\n", c.FindHardwareAddress())
}
