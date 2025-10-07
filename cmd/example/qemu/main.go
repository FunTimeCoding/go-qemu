package main

import "github.com/funtimecoding/go-qemu/pkg/qemu/example"

func main() {
	example.Version()

	if false {
		example.Domain()
		example.CreateDisk()
		example.CreateDomain()
		example.DeleteDomain()
		example.HardwareAddress()
		example.Wrapper()
	}
}
