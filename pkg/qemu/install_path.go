package qemu

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"path"
)

func InstallPath(binary string) string {
	p := system.LookPath(binary)

	return path.Clean(system.Join(path.Dir(p), system.ReadLink(p), "..", ".."))
}
