package apis

import (
	"github.com/go-jarvis/rum-gonic/rum"
	"github.com/tangx/srv-go-blog/cmd/srv-go-blog/apis/index"
)

var (
	RouterGroup_Root = rum.NewRouterGroup("/")
	RouterGroup_V0   = rum.NewRouterGroup("/0")
)

func init() {
	RouterGroup_Root.Register(RouterGroup_V0)

	RouterGroup_V0.Register(&index.Index{})
}
