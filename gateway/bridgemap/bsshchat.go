// +build !nosshchat

package bridgemap

import (
	bsshchat "github.com/yujintong/matterbridge/bridge/sshchat"
)

func init() {
	FullMap["sshchat"] = bsshchat.New
}
