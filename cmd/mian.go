package main

import (
	"NotMe/asset"
	"NotMe/clipboard"
	"NotMe/utils"

	_ "NotMe/asset"
	_ "NotMe/config"
	"fmt"
)

func main() {
	for {
		ips := <-clipboard.Chan
		fmt.Println(len(ips))
		isTure, msg := asset.CheckMany(ips)
		if isTure {
			err := utils.Notify(msg)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
