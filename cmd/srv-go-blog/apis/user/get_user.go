package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/tangx/srv-go-blog/pkg/daos"
)

func init() {
	RouterGroup_User.Register(&GetUserByID{})
}

type GetUserByID struct {
	httpx.MethodGet `path:"/:id"`
	ID              int `uri:"id"`
}

func (u *GetUserByID) Output(c *gin.Context) (interface{}, error) {

	author, err := daos.GetUserByID(c, u.ID)
	if err != nil {
		return nil, err
	}
	return author, nil

}
