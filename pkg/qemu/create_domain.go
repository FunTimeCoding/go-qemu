package qemu

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
	"net"
)

const (
	Gibibyte = "GiB"

	SmallType      = "scsi"
	PeripheralType = "pci"

	ExpressRoot     = "pcie-root"
	ExpressRootPort = "pcie-root-port"
)

func (c *Client) CreateDomain(
	name string,
	cores uint,
	memoryGibibytes uint,
	h net.HardwareAddr,
	installImage string,
) *libvirt.Domain {
	ports := uint(15)
	// noinspection HttpUrlsUsage
	var d = &libvirtxml.Domain{
		Name: name,
		Type: "hvf",
		CPU: &libvirtxml.DomainCPU{
			Mode:  "custom",
			Match: "exact",
			Model: &libvirtxml.DomainCPUModel{Value: "cortex-a57"},
		},
		Metadata: &libvirtxml.DomainMetadata{
			XML: `
    <libosinfo:libosinfo xmlns:libosinfo="http://libosinfo.org/xmlns/libvirt/domain/1.0">
        <libosinfo:os id="http://debian.org/debian/12"/>
    </libosinfo:libosinfo>
  `,
		},
		Memory: &libvirtxml.DomainMemory{
			Value: memoryGibibytes,
			Unit:  Gibibyte,
		},
		CurrentMemory: &libvirtxml.DomainCurrentMemory{
			Value: memoryGibibytes,
			Unit:  Gibibyte,
		},
		VCPU: &libvirtxml.DomainVCPU{Value: cores},
		OS: &libvirtxml.DomainOS{
			Firmware: "efi",
			Type: &libvirtxml.DomainOSType{
				Arch:    "aarch64",
				Machine: "virt",
				Type:    "hvm",
			},
			BootDevices: []libvirtxml.DomainBootDevice{
				{Dev: "cdrom"},
				{Dev: "hd"},
			},
		},
		Features: &libvirtxml.DomainFeatureList{
			ACPI: &libvirtxml.DomainFeature{},
		},
		Clock: &libvirtxml.DomainClock{
			Offset: "utc", // localtime
		},
		Devices: &libvirtxml.DomainDeviceList{
			Emulator: "/opt/homebrew/bin/qemu-system-aarch64",
			Disks: []libvirtxml.DomainDisk{
				{
					Device: "disk",
					Driver: &libvirtxml.DomainDiskDriver{
						Name:    "qemu",
						Type:    "qcow2",
						Discard: "unmap",
					},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: system.Join(
								c.diskDirectory,
								fmt.Sprintf("%s.qcow2", name),
							),
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "vda",
						Bus: "virtio",
					},
				},
				{
					Device: "cdrom",
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
						Type: "raw",
					},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: installImage,
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "sda",
						Bus: "scsi",
					},
					ReadOnly: &libvirtxml.DomainDiskReadOnly{},
				},
			},
			Controllers: []libvirtxml.DomainController{
				{
					Type: "usb", Model: "qemu-xhci",
					USB: &libvirtxml.DomainControllerUSB{Port: &ports},
				},
				{Type: SmallType, Model: "virtio-scsi"},
				{Type: PeripheralType, Model: ExpressRoot},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
				{Type: PeripheralType, Model: ExpressRootPort},
			},
			Interfaces: []libvirtxml.DomainInterface{
				{
					MAC: &libvirtxml.DomainInterfaceMAC{
						Address: h.String(),
					},
					Model: &libvirtxml.DomainInterfaceModel{Type: "virtio"},
				},
			},
			Consoles: []libvirtxml.DomainConsole{
				{
					Source: &libvirtxml.DomainChardevSource{
						Pty: &libvirtxml.DomainChardevSourcePty{},
					},
				},
			},
			Channels: []libvirtxml.DomainChannel{
				{
					Source: &libvirtxml.DomainChardevSource{
						UNIX: &libvirtxml.DomainChardevSourceUNIX{Mode: "bind"},
					},
					Target: &libvirtxml.DomainChannelTarget{
						VirtIO: &libvirtxml.DomainChannelTargetVirtIO{
							Name: "org.qemu.guest_agent.0",
						},
					},
				},
			},
			MemBalloon: &libvirtxml.DomainMemBalloon{Model: "virtio"},
			RNGs: []libvirtxml.DomainRNG{
				{
					Model: "virtio",
					Backend: &libvirtxml.DomainRNGBackend{
						Random: &libvirtxml.DomainRNGBackendRandom{
							Device: "/dev/urandom",
						},
					},
				},
			},
		},
	}
	create, marshalFail := d.Marshal()
	errors.PanicOnError(marshalFail)
	create = CollapseEmptyTags(create)
	fmt.Printf("Create: %s\n", create)
	result, e := c.virt.DomainCreateXML(create, 0)
	errors.PanicOnError(e)

	return result
}
