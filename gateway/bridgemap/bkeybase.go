// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/yujintong/matterbridge/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}
