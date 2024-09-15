package kata6


import (
	"strconv"
	"strings"
	"net"
)

/*
#IP Validation
Write an algorithm that will identify valid IPv4 addresses in dot-decimal format. IPs should be considered valid if they consist of four octets, with values between 0 and 255, inclusive.

Valid inputs examples:
Examples of valid inputs:
1.2.3.4
123.45.67.89
Invalid input examples:
1.2.3
1.2.3.4.5
123.456.78.90
123.045.067.089
Notes:
Leading zeros (e.g. 01.02.03.04) are considered invalid
Inputs are guaranteed to be a single string
*/
func IsValidIp2(ip string) bool {
	if len(ip) == 0 {
		return false
	}
	if !strings.ContainsRune(ip, '.') {
		return false
	}
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}
	for _,ch :=range parts {		
		if len(ch) > 1 && string(ch[0]) == "0" {
			return false
		}
		n, err:= strconv.Atoi(ch)
		if err != nil {
			return false
		}
		if n<0 || n> 255 {
			return false
		}
				
	}

	return true
}

func IsValidIp(ip string) bool {
	res := net.ParseIP(ip)
	if res == nil {
		return false
	}
	return true
}