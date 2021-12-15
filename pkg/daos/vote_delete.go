package daos

import (
	"context"

	"github.com/tangx/srv-go-blog/pkg/injectors/db"
	"github.com/tangx/srv-go-blog/pkg/models"
)

func DeleteVoteByID(ctx context.Context, id uint) error {

	d := db.FromInjectedContext(ctx)

	t := d.Delete(&models.Vote{}, id)

	if t.Error != nil {
		return t.Error
	}

	return nil
}
