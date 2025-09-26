package main

import (
	"NotMe/clipboard"
	"fmt"
)

func main() {
	for {
		ip := <-clipboard.Chan
		fmt.Println(len(ip))
		for _, a := range ip {
			fmt.Println(a)
		}
	}

}
