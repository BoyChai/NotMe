package asset

import (
	"NotMe/config"
	"NotMe/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadText(cfg []config.Text) {
	utils := utils.Asset
	for _, c := range cfg {
		file, err := os.Open(c.Path)
		if err != nil {
			fmt.Println("打开文本失败:", err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}

			// 如果有分隔符，就切分
			var ipOrCidr, desc string
			if c.Delim != "" {
				parts := strings.Split(line, c.Delim)
				if len(parts) >= 2 {
					ipOrCidr = strings.TrimSpace(parts[0])
					desc = strings.TrimSpace(parts[1])
				}
			} else {
				ipOrCidr = line
			}

			if utils.IsValidIPv4(ipOrCidr) {
				if ok := addIP(ipOrCidr, desc); !ok {
					fmt.Println("添加IP失败:", ipOrCidr)
				}
				continue
			}
			if utils.IsValidCIDR(ipOrCidr) {
				if ok := addCIDR(ipOrCidr, desc); !ok {
					fmt.Println("添加CIDR失败:", ipOrCidr)
				}
				continue
			}
			// TODO: 未识别的处理逻辑
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("读取文本出错:", err)
		}
	}
}
