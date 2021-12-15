package post

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/tangx/srv-go-blog/pkg/daos"
)

func init() {
	RouterGroupPost.Register(&GetPostByID{})
}

type GetPostByID struct {
	httpx.MethodGet `path:"/:id"`
	ID              uint `uri:"id"`
}

func (req *GetPostByID) Output(c *gin.Context) (interface{}, error) {
	post, err := daos.GetPostByID(c, req.ID)
	if err != nil {
		return nil, err
	}
	return post, nil
}
