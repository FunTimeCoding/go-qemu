package hardware_address

import "net"

func FirstUnused(
	start net.HardwareAddr,
	v []net.HardwareAddr,
) net.HardwareAddr {
	result := start

	for _, e := range v {
		if LessThan(result, e) {
			break
		}

		result = Increment(result)
	}

	return result
}
