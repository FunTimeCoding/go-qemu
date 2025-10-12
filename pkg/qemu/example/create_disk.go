package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"github.com/funtimecoding/go-qemu/pkg/qemu"
)

func CreateDisk() {
	c := qemu.New()
	defer c.Close()

	diskPath := join.Absolute(
		system.Home(),
		constant.DownloadsPath,
		fmt.Sprintf("%s.qcow2", c.Domain()),
	)

	if !system.FileExists(diskPath) {
		fmt.Printf("Disk: %+v\n", c.CreateDisk(diskPath, 12))
	} else {
		fmt.Printf("Disk exists: %s\n", diskPath)
	}
}
