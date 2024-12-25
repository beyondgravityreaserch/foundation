package validation

import (
	"net"
)

// IP function checks if the given IP address string is an IPv4 or IPv6 address.
//
// Example:
//
//	fmt.Println(IP("192.168.1.1")) // Output: true
//	fmt.Println(IP("::1"))         // Output: false
func IP(ip_str string) bool {
	ip := net.ParseIP(ip_str)
	if ip.To4() != nil {
		return true // IPv4
	} else {
		return false // IPv6
	}
}
