package global

import (
	"github.com/go-jarvis/confgorm/drivers"
	"github.com/go-jarvis/confgorm/migration"
	"github.com/go-jarvis/confhttp"
	"github.com/go-jarvis/jarvis"
	"github.com/go-jarvis/jarvis/pkg/appctx"
	"github.com/tangx/srv-go-blog/cmd/srv-go-blog/apis"
	"github.com/tangx/srv-go-blog/pkg/injectors/db"
	"github.com/tangx/srv-go-blog/pkg/models"
)

var (
	httpServer = &confhttp.Server{}
	masterDB   = &drivers.MysqlDriver{
		MigrationDB: models.DB,
	}
)

var (
	App = jarvis.New().WithOptions(
		appctx.WithName("srv-go-blog"),
		appctx.WithRoot("../.."),
	)
)

func init() {
	config := &struct {
		HttpServer *confhttp.Server
		MasterDB   *drivers.MysqlDriver
	}{
		HttpServer: httpServer,
		MasterDB:   masterDB,
	}

	_ = App.Conf(config)
}

func Server() *confhttp.Server {

	httpServer.WithContextInjectors(
		db.InjectDBInto(masterDB),
	)

	httpServer.Register(apis.RouterGroup_Root)
	return httpServer
}

func AutoMigrate() {
	migration.Magrate(masterDB)
}
