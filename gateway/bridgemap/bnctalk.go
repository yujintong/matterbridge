// +build !nonctalk

package bridgemap

import (
	btalk "github.com/yujintong/matterbridge/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}
