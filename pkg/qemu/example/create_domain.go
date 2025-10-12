package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/debian"
	"github.com/funtimecoding/go-library/pkg/debian/image"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"github.com/funtimecoding/go-qemu/pkg/qemu"
)

func CreateDomain() {
	c := qemu.New()
	defer c.Close()

	v := debian.Bookworm.Version()
	installImage := join.Absolute(
		system.Home(),
		constant.DownloadsPath,
		image.Name(v, "arm64"),
	)
	fmt.Printf(
		"Domain: %+v",
		c.CreateDomain(
			c.Domain(),
			1,
			1,
			c.FindHardwareAddress(),
			installImage,
		),
	)
}
