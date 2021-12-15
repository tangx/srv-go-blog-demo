package daos

import (
	"context"

	"github.com/go-jarvis/statuserrors"
	"github.com/tangx/srv-go-blog/pkg/injectors/db"
	"github.com/tangx/srv-go-blog/pkg/models"
	"gorm.io/gorm"
)

func VotePost(ctx context.Context, postID, authorID uint, status string) error {
	d := db.FromInjectedContext(ctx)

	vote := &models.Vote{}

	t := d.Where("author_id=? and post_id = ?", authorID, postID).First(vote)

	if t.Error != nil && statuserrors.Is(t.Error, gorm.ErrRecordNotFound) {
		//todo: create
		vote.VoteInfo = models.VoteInfo{
			AuthorID:  authorID,
			PostID:    postID,
			LikeOrNot: status,
		}

		t := d.Create(vote)
		return t.Error
	}

	// todo: update
	vote.VoteInfo.LikeOrNot = status
	t = d.Model(vote).Select("like_or_not").Updates(vote)

	return t.Error

}
