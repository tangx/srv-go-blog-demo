package daos

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("user not found: %d", id)
	}

	if t.Error != nil {
		return nil, statuserrors.Wrap(
			t.Error,
			statuserrors.StatusUnknownError,
			"unknown error: %v", t.Error,
		)
	}

	// fmt.Println("t.RowsAffected=", t.RowsAffected)

	return user, nil

}
