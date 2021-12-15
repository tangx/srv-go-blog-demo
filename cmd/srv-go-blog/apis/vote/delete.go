package vote

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/tangx/srv-go-blog/pkg/daos"
)

func init() {
	RouterGroupVote.Register(&DeleteVoteByID{})
}

type DeleteVoteByID struct {
	httpx.MethodDelete `path:"/id/:id"`
	VoteID             uint `uri:"id"`
}

func (req *DeleteVoteByID) Output(c *gin.Context) (interface{}, error) {
	err := daos.DeleteVoteByID(c, req.VoteID)

	if err != nil {
		return nil, err
	}

	return "delete success", nil
}
