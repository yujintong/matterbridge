// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/yujintong/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
