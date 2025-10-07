package example

import (
	"fmt"
	upstream "github.com/digitalocean/go-qemu/qemu"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-qemu/pkg/qemu"
)

func Wrapper() {
	c := qemu.New()
	defer c.Close()

	m := c.Monitor(c.Domain())

	domain, domainFail := upstream.NewDomain(m, m.Domain)
	errors.PanicOnError(domainFail)
	defer func() {
		errors.LogClose(domain)
	}()

	version, versionFail := domain.Version()
	errors.PanicOnError(versionFail)
	fmt.Printf("Version: %s\n", version)

	status, statusFail := domain.Status()
	errors.PanicOnError(statusFail)
	fmt.Printf("Status: %s\n", status)
}
