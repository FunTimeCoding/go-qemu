package hardware_address

import "net"

func LessThan(
	a net.HardwareAddr,
	b net.HardwareAddr,
) bool {
	count := len(a)

	for i := 0; i < count; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}

	return false
}
