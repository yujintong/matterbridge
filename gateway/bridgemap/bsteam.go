// +build !nosteam

package bridgemap

import (
	bsteam "github.com/yujintong/matterbridge/bridge/steam"
)

func init() {
	FullMap["steam"] = bsteam.New
}
