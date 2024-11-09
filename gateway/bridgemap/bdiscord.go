// +build !nodiscord

package bridgemap

import (
	bdiscord "github.com/yujintong/matterbridge/bridge/discord"
)

func init() {
	FullMap["discord"] = bdiscord.New
	UserTypingSupport["discord"] = struct{}{}
}
