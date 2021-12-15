package daos

import (
	"context"

	"github.com/tangx/srv-go-blog/pkg/injectors/db"
	"github.com/tangx/srv-go-blog/pkg/models"
	"gorm.io/gorm"
)

// IncreasePostReadingCount reading count +1
// https://gorm.io/docs/update.html#Update-with-SQL-Expression
func IncreasePostReadingCount(ctx context.Context, id uint) {
	d := db.FromInjectedContext(ctx)

	d.Model(&models.Post{}).
		Where("id = ?", id).
		Update(
			"reading_count",
			gorm.Expr("reading_count + 1"),
		)
}
