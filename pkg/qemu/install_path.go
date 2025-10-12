package qemu

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"path"
)

func InstallPath(binary string) string {
	p := system.LookPath(binary)

	return path.Clean(
		join.Absolute(path.Dir(p), system.ReadLink(p), "..", ".."),
	)
}
