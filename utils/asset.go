package utils

import (
	"net"
	"strings"
)

type asset struct{}

var Asset asset

func (asset) ColLetterToIndex(col string) int {
	col = strings.ToUpper(col)
	index := 0
	for i := 0; i < len(col); i++ {
		index = index*26 + int(col[i]-'A'+1)
	}
	return index - 1
}

// 判断是否为合法 IPv4
func (asset) IsValidIPv4(ip string) bool {
	parsed := net.ParseIP(strings.TrimSpace(ip))
	return parsed != nil && parsed.To4() != nil
}
func (asset) IsValidCIDR(cidr string) bool {
	_, _, err := net.ParseCIDR(strings.TrimSpace(cidr))
	return err == nil
}
