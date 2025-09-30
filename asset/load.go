package asset

import (
	"NotMe/config"
	"fmt"
	"net"
	"strings"
)

type CIDR struct {
	CIDR    net.IPNet // 网段
	Comment string    // 备注
}

type assetStore struct {
	IPs   map[string]string // 纯 IP 地址
	CIDRs []CIDR            // 网段
}

var Store assetStore = assetStore{
	IPs: make(map[string]string),
}

// 加载资产
func init() {
	var cfg = config.Global
	loadXlsx(cfg.XLSX)
	loadText(cfg.TEXT)
	fmt.Printf("总计加载 %d IPs 和 %d CIDRs\n", len(Store.IPs), len(Store.CIDRs))
}

func addIP(ip, comment string) bool {
	Store.IPs[ip] = comment
	return true
}

func addCIDR(cidr, comment string) bool {
	_, ipnet, err := net.ParseCIDR(strings.TrimSpace(cidr))
	if err != nil {
		fmt.Println(err)
		return false
	}
	Store.CIDRs = append(Store.CIDRs, CIDR{
		CIDR:    *ipnet,
		Comment: comment,
	})
	return true
}
