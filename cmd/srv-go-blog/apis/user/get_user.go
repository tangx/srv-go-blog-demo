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
	httpx.MethodGet `path:"/id/:id"`
	ID              int `uri:"id"`
}

func (u *GetUserByID) Output(c *gin.Context) (interface{}, error) {

	author, err := daos.GetUserByID(c, u.ID)
	if err != nil {
		return nil, err
	}
	return author, nil

}

type GetUserByName struct {
	httpx.MethodGet `path:"/name/:name"`
	Name            string `uri:"name"`
}

func (index *GetUserByName) Output(c *gin.Context) (interface{}, error) {
	author, err := daos.GetUserByName(c, index.Name)
	if err != nil {
		return nil, err
	}

	return author, nil
}
