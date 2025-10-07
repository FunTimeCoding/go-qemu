package qemu

import (
	digital "github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-libvirt/socket/dialers"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"libvirt.org/go/libvirt"
	"time"
)

func New() *Client {
	virt, f := libvirt.NewConnect("qemu:///system")
	errors.PanicOnError(f)

	socket := "/opt/homebrew/var/run/libvirt/libvirt-sock"
	timeout := 2 * time.Second
	d := digital.NewWithDialer(
		dialers.NewLocal(
			dialers.WithSocket(socket),
			dialers.WithLocalTimeout(timeout),
		),
	)
	errors.PanicOnError(d.Connect())

	return &Client{
		virt:            virt,
		digital:         d,
		timeout:         timeout,
		socket:          socket,
		domain:          "development",
		hardwareAddress: "52:54:00:00:00:00",
		diskDirectory:   system.Join(system.Home(), constant.DownloadsPath),
	}
}
