// +build !nomatrix

package bridgemap

import (
	bmatrix "github.com/yujintong/matterbridge/bridge/matrix"
)

func init() {
	FullMap["matrix"] = bmatrix.New
}
