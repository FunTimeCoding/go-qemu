package qemu

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
)

func (c *Client) CreateDisk(
	path string,
	size int,
) string {
	// Formatting '/Users/shiin/Downloads/development.qcow2', fmt=qcow2 cluster_size=65536 extended_l2=off compression_type=zlib size=12884901888 lazy_refcounts=off refcount_bits=16
	return system.Run(
		"qemu-img",
		"create",
		"-f",
		"qcow2",
		path,
		fmt.Sprintf("%dG", size),
	)
}
