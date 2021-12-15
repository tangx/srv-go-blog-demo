package daos

import (
	"context"
	"net/http"

	"github.com/go-jarvis/statuserrors"
	"github.com/tangx/srv-go-blog/pkg/injectors/db"
	"github.com/tangx/srv-go-blog/pkg/models"
	"gorm.io/gorm"
)

func GetPostByID(ctx context.Context, id uint) (*models.Post, error) {
	d := db.FromInjectedContext(ctx)

	post := &models.Post{}
	t := d.First(post, id)
	if t.Error != nil && statuserrors.Is(t.Error, gorm.ErrRecordNotFound) {

		return nil, statuserrors.Wrap(t.Error, http.StatusNotFound, "post %d not found", id)

	}

	if t.Error != nil {
		return nil, t.Error
	}

	go IncreasePostReadingCount(ctx, id)

	return post, nil
}
