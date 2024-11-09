// +build !nomumble

package bridgemap

import (
	bmumble "github.com/yujintong/matterbridge/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}
