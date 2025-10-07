package qemu

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) Version() int {
	result, e := c.virt.GetLibVersion()
	errors.PanicOnError(e)

	return int(result)
}
