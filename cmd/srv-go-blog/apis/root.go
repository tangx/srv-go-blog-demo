package apis

import (
	"github.com/go-jarvis/rum-gonic/rum"
	"github.com/tangx/srv-go-blog/cmd/srv-go-blog/apis/index"
	"github.com/tangx/srv-go-blog/cmd/srv-go-blog/apis/post"
	"github.com/tangx/srv-go-blog/cmd/srv-go-blog/apis/user"
)

var (
	RouterGroup_Root = rum.NewRouterGroup("/")
	RouterGroup_V0   = rum.NewRouterGroup("/v0")
)

func init() {
	RouterGroup_Root.Register(RouterGroup_V0)

	{
		RouterGroup_V0.Register(&index.Index{})
		RouterGroup_V0.Register(user.RouterGroup_User)

		RouterGroup_V0.Register(post.RouterGroupPost)
	}

}
