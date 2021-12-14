package db

import (
	"context"

	"github.com/go-jarvis/confgorm/drivers/mysql"

	"github.com/go-jarvis/confgorm/drivers"

	"github.com/go-jarvis/rum-gonic/rum"
)

type Database string

const (
	KEY_MasterDB Database = "KEY_MasterDB"
)

func InjectDBInto(driver *drivers.MysqlDriver) rum.ContextInjectorFunc {
	return rum.InjectContextValueWith(KEY_MasterDB, driver)
}

func FromInjectedContext(ctx context.Context) *mysql.MysqlDriver {
	obj := rum.InjectedContextValueFrom(ctx, KEY_MasterDB)
	return obj.(*mysql.MysqlDriver)
}
