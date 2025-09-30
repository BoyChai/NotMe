package clipboard

import (
	"NotMe/utils"
	"time"

	"github.com/atotto/clipboard"
)

var Chan = make(chan []string)

var lastContent string

func init() {

	utils := utils.Clipboard

	pollInterval := 500 * time.Millisecond
	go func() {
		for {
			currentContent, err := clipboard.ReadAll()
			if err != nil || currentContent == "" || currentContent == lastContent {
				time.Sleep(pollInterval)
				continue
			}
			lastContent = currentContent

			if !utils.IsProbablyText(currentContent) {
				continue
			}

			lastContent = currentContent
			data := utils.SplitByNewline(currentContent)
			var list []string

			for _, line := range data {
				if utils.IsValidIPv4(line) {
					list = append(list, line)
					continue
				}
			}
			Chan <- list
			time.Sleep(pollInterval)
		}
	}()

}
