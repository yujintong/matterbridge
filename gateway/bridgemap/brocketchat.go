// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/yujintong/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
