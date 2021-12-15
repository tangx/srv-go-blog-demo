package daos

import (
	"context"

	"github.com/tangx/srv-go-blog/pkg/injectors/db"
	"github.com/tangx/srv-go-blog/pkg/models"
)

func CreateAuthor(ctx context.Context, name, password string) (*models.Author, error) {
	d := db.FromInjectedContext(ctx)

	user := &models.Author{
		Name:     name,
		Password: password,
	}

	tx := d.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user.Desensitization(), nil
}
