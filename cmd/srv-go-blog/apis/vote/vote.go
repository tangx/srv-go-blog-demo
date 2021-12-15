package vote

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/tangx/srv-go-blog/pkg/daos"
)

func init() {
	RouterGroupVote.Register(&VoteByPostID{})
}

type VoteByPostID struct {
	httpx.MethodPost `path:"/post/:id"`
	AuthorName       string `header:"AuthorName"`
	PostID           uint   `uri:"id"`
	Data             struct {
		Status string `json:"status"`
	} `body:"body"`
}

func (req *VoteByPostID) Output(c *gin.Context) (interface{}, error) {

	user, err := daos.GetUserByName(c, req.AuthorName)
	if err != nil {
		return nil, err
	}

	err = daos.VotePost(c, req.PostID, user.ID, req.Data.Status)
	if err != nil {
		return nil, err
	}

	return "vote success", nil
}

// type VoteStatus string

// const (
// 	VoteLike    VoteStatus = "like"
// 	VoteDislike VoteStatus = "dislike"
// 	VoteNothing VoteStatus = "nothing"
// )
