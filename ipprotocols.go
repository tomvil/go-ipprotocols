package ipprotocols

import (
	"fmt"
	"strings"
)

func GetProtocolName(protocol int) (string, error) {
	if name, exists := protocolNumbers[protocol]; exists {
		return name, nil
	}

	return "Unknown", fmt.Errorf("unknown protocol: %d", protocol)
}

func GetProtocolNumber(protocol string) (int, error) {
	protocol = strings.ToLower(protocol)
	if number, exists := protocolNames[protocol]; exists {
		return number, nil
	}

	return -1, fmt.Errorf("unknown protocol: %s", protocol)
}
