package qemu

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) Close() {
	if c.virt != nil {
		_, e := c.virt.Close()
		errors.PanicOnError(e)
	}

	if c.digital != nil {
		errors.PanicOnError(c.digital.Disconnect())
	}

	for _, m := range c.monitors {
		errors.PanicOnError(m.Disconnect())
	}
}
