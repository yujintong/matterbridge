// +build !noirc

package bridgemap

import (
	birc "github.com/yujintong/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
