package walllet

import (
	"fmt"
	"strings"
)

func ValidateAddress(address string) error {

	if address == "" {
		return fmt.Errorf("address is empty")
	}

	if !strings.HasPrefix(address, "0x") {
		return fmt.Errorf("address must start with 0x")
	}

	if len(address) != 42 {
		return fmt.Errorf("address length must be 42")
	}

	for _, ch := range address[2:] {
		if !isHexChar(ch) {
			return fmt.Errorf("address contains non-hex character :%c", ch)
		}
	}
	return nil
}

func isHexChar(ch rune) bool {

	return ch >= '0' && ch <= '9' ||
		ch >= 'a' && ch <= 'f' ||
		ch >= 'A' && ch <= 'F'
}
