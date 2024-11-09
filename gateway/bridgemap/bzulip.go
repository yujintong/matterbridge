// +build !nozulip

package bridgemap

import (
	bzulip "github.com/yujintong/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
