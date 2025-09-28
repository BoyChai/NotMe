package asset

import (
	"fmt"
	"net"
	"strings"
)

func CheckMany(ips []string) (bool, string) {
	var results []string

	for _, ip := range ips {
		ip = strings.TrimSpace(ip)
		parsed := net.ParseIP(ip)
		if parsed == nil {
			continue
		}

		if comment, ok := Store.IPs[ip]; ok {
			results = append(results, fmt.Sprintf("%s -> %s", ip, comment))
		}

		for _, c := range Store.CIDRs {
			if c.CIDR.Contains(parsed) {
				results = append(results, fmt.Sprintf("%s -> %s", c.CIDR.String(), c.Comment))
			}
		}
	}

	if len(results) == 0 {
		return false, ""
	}

	// 2条展示，其他汇总
	if len(results) > 2 {
		display := strings.Join(results[:2], "\n")
		return true, fmt.Sprintf("%s\n共 %d 个匹配记录", display, len(results))
	}

	return true, strings.Join(results, "\n")
}
