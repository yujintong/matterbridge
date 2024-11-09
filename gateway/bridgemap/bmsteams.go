// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/yujintong/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
