package example

import (
	"fmt"
	"github.com/funtimecoding/go-qemu/pkg/qemu"
)

func Version() {
	c := qemu.New()
	defer c.Close()
	fmt.Printf("Version: %d\n", c.Version())
}
