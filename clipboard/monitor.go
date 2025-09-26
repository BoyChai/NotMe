package clipboard

import (
	"time"

	"github.com/atotto/clipboard"
)

var Chan = make(chan []string)

var lastContent string

func init() {
	pollInterval := 500 * time.Millisecond
	go func() {
		for {
			currentContent, err := clipboard.ReadAll()
			if err != nil {
				continue
			}
			if currentContent == lastContent || currentContent == "" {
				continue
			}
			lastContent = currentContent
			data := splitByNewline(currentContent)
			var list []string

			for _, line := range data {
				if isValidIPv4(line) {
					list = append(list, line)
				}
			}
			Chan <- list
			time.Sleep(pollInterval)
		}
	}()

}
