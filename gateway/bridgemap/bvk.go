// +build !novk

package bridgemap

import (
	bvk "github.com/yujintong/matterbridge/bridge/vk"
)

func init() {
	FullMap["vk"] = bvk.New
}
