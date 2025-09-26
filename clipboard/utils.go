package clipboard

import (
	"strconv"
	"strings"
)

// IsValidIPv4 判断字符串是否是有效的 IPv4 地址（不包括掩码）
func isValidIPv4(ip string) bool {
	ip = strings.ReplaceAll(ip, "\r", "")

	if strings.Contains(ip, "/") {
		parts := strings.Split(ip, "/")
		ip = parts[0]
	}

	segments := strings.Split(ip, ".")
	if len(segments) != 4 {
		return false
	}

	for _, segment := range segments {
		if len(segment) == 0 || (len(segment) > 1 && segment[0] == '0') {
			return false
		}

		num, err := strconv.Atoi(segment)
		if err != nil {

			return false
		}

		if num < 0 || num > 255 {
			return false
		}
	}

	return true
}

func splitByNewline(input string) []string {
	if input == "" {
		return []string{}
	}
	lines := strings.Split(input, "\n")
	if len(lines) == 1 && lines[0] == input {
		return []string{input}
	}
	result := []string{}
	for _, line := range lines {
		if line != "" {
			result = append(result, line)
		}
	}
	return result
}
