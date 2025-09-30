package main

import (
	"NotMe/asset"
	_ "NotMe/asset"
	"NotMe/clipboard"
	_ "NotMe/config"
	"NotMe/utils"
	"fmt"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(utils.Icon)
	systray.SetTitle("NotMe")
	systray.SetTooltip("NotMe")
	quit := systray.AddMenuItem("退出", "退出程序")
	go func() {
		<-quit.ClickedCh
		systray.Quit()
	}()

	for {
		ips := <-clipboard.Chan
		// fmt.Println(len(ips))
		isTure, msg := asset.CheckMany(ips)
		if isTure {
			err := utils.Notify(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func onExit() {
	close(clipboard.Chan)
}
