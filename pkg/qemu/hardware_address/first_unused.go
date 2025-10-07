package hardware_address

import "net"

func FirstUnused(
	start net.HardwareAddr,
	input []net.HardwareAddr,
) net.HardwareAddr {
	result := start

	for _, element := range input {
		if LessThan(result, element) {
			break
		}

		result = Increment(result)
	}

	return result
}
