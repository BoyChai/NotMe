package utils

import (
	_ "embed"

	"github.com/gen2brain/beeep"
)

const title = "NotMe"

//go:embed static/icon.png
var icon []byte

// Beeep支持跨平台，但是有局限性无法创建点击事件
// 后期对于每个平台单独做处理，支持事件详情
func Notify(msg string) error {
	beeep.AppName = "NotMe"
	title := "资产匹配告警"
	return beeep.Alert(title, msg, icon)
}
