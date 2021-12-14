package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-go-blog/pkg/daos"
)

func init() {
	RouterGroup_User.Register(&CreateUser{})
}

type CreateUser struct {
	httpx.MethodPost `path:""`

	Data CreateUserParams `body:"body"`
}

type CreateUserParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *CreateUser) Output(c *gin.Context) (interface{}, error) {
	user, err := daos.CreateAuthor(c, u.Data.Name, u.Data.Password)
	if err != nil {
		logrus.Infof("create user failed: %v", err)
		return nil, err
	}
	return user, nil
}
