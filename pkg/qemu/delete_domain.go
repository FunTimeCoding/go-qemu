package qemu

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) DeleteDomain(name string) {
	for _, d := range c.domains() {
		if Name(&d) == name {
			if IsPersistent(&d) {
				errors.PanicOnError(d.Undefine())
			}

			errors.PanicOnError(d.Destroy())
			errors.PanicOnError(d.Free())

			break
		}
	}
}
