package hardware_address

import "net"

func Increment(a net.HardwareAddr) net.HardwareAddr {
	result := make(net.HardwareAddr, len(a))
	copy(result, a)

	for i := len(result) - 1; i >= 0; i-- {
		result[i]++

		if result[i] != 0 {
			break
		}
	}

	return result
}
