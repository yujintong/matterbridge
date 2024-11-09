// +build whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/yujintong/matterbridge/bridge/whatsappmulti"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
