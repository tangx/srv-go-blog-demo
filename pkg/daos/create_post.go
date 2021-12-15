package daos

import (
	"context"
	"net/http"

	"github.com/go-jarvis/statuserrors"
	"github.com/tangx/srv-go-blog/pkg/injectors/db"
	"github.com/tangx/srv-go-blog/pkg/models"
)

func CreatePost(ctx context.Context, title, content string, authorID uint) (*models.Post, error) {
	d := db.FromInjectedContext(ctx)

	post := &models.Post{
		Title:    title,
		Content:  content,
		AuthorID: authorID,
	}

	t := d.Create(post)
	if t.Error != nil {
		return nil, statuserrors.Wrap(t.Error, http.StatusInternalServerError, "create post failed: %s", title)
	}

	return post, nil

}
