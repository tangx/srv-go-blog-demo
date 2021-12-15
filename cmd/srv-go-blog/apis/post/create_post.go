package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/statuserrors"
	"github.com/tangx/srv-go-blog/pkg/daos"
)

func init() {
	RouterGroupPost.Register(&CreatePost{})
}

type CreatePost struct {
	httpx.MethodPost `path:""`
	AuthorName       string `header:"AuthorName"`
	Data             struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	} `body:"body"`
}

func (req *CreatePost) Output(c *gin.Context) (interface{}, error) {
	author, err := daos.GetUserByName(c, req.AuthorName)
	if err != nil {
		return nil, statuserrors.New(http.StatusBadRequest, "invalid user")
	}

	post, err := daos.CreatePost(c, req.Data.Title, req.Data.Content, author.ID)
	if err != nil {
		return nil, err
	}

	return post, nil
}
