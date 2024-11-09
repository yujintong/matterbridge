// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/yujintong/matterbridge/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}
