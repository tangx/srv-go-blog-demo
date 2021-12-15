package daos

import (
	"context"
	"net/http"

	"github.com/go-jarvis/statuserrors"
	"github.com/tangx/srv-go-blog/pkg/injectors/db"
	"github.com/tangx/srv-go-blog/pkg/models"
	"gorm.io/gorm"
)

func GetUserByID(ctx context.Context, id int) (*models.Author, error) {
	d := db.FromInjectedContext(ctx)

	user := &models.Author{}
	t := d.First(user, id)

	if statuserrors.Is(t.Error, gorm.ErrRecordNotFound) {
		return nil, statuserrors.Wrap(t.Error, http.StatusNotFound, "user not found: %d", id)
	}

	if t.Error != nil {
		return nil, statuserrors.Wrap(
			t.Error,
			statuserrors.StatusUnknownError,
			"unknown error: %v", t.Error,
		)
	}

	return user.Desensitization(), nil

}

func GetUserByName(ctx context.Context, name string) (*models.Author, error) {
	d := db.FromInjectedContext(ctx)

	user := &models.Author{}
	t := d.Where("name = ?", name).First(user)
	if statuserrors.Is(t.Error, gorm.ErrRecordNotFound) {
		return nil, statuserrors.Wrap(t.Error, http.StatusNotFound, "user not found: %s", name)
	}

	if t.Error != nil {
		return nil, t.Error
	}

	return user.Desensitization(), nil
}
