// +build !notelegram

package bridgemap

import (
	btelegram "github.com/yujintong/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
