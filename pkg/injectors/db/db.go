package db

import (
	"context"

	"gorm.io/gorm"

	"github.com/go-jarvis/rum-gonic/rum"
)

type Database string

const (
	KEY_MasterDB Database = "KEY_MasterDB"
)

func InjectDBInto(db *gorm.DB) rum.ContextInjectorFunc {
	return rum.InjectContextValueWith(KEY_MasterDB, db)
}

func FromInjectedContext(ctx context.Context) *gorm.DB {
	obj := rum.InjectedContextValueFrom(ctx, KEY_MasterDB)
	return obj.(*gorm.DB)
}
