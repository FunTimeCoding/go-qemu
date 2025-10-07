package qemu

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) DeleteDomain(name string) {
	for _, element := range c.domains() {
		if Name(&element) == name {
			if IsPersistent(&element) {
				errors.PanicOnError(element.Undefine())
			}

			errors.PanicOnError(element.Destroy())
			errors.PanicOnError(element.Free())

			break
		}
	}
}
