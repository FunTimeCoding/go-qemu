package qemu

import (
	digital "github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-qemu/qmp"
	"libvirt.org/go/libvirt"
	"time"
)

type Client struct {
	virt            *libvirt.Connect
	digital         *digital.Libvirt
	monitors        []*qmp.LibvirtRPCMonitor
	socket          string
	timeout         time.Duration
	domain          string
	diskDirectory   string
	hardwareAddress string
}
