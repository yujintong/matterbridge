// +build !noapi

package bridgemap

import (
	"github.com/yujintong/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
