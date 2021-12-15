package global

import (
	"github.com/go-jarvis/confgorm/drivers/mysqldriver"
	"github.com/go-jarvis/confgorm/migration"
	"github.com/go-jarvis/jarvis"
	"github.com/go-jarvis/jarvis/pkg/appctx"
	"github.com/go-jarvis/rum-gonic/confhttp"
	"github.com/tangx/srv-go-blog/cmd/srv-go-blog/apis"
	"github.com/tangx/srv-go-blog/pkg/injectors/db"
	"github.com/tangx/srv-go-blog/pkg/models"
)

var (
	httpServer = &confhttp.Server{}
	masterDB   = &mysqldriver.Server{
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
		MasterDB   *mysqldriver.Server
	}{
		HttpServer: httpServer,
		MasterDB:   masterDB,
	}

	_ = App.Conf(config)
}

func Server() *confhttp.Server {

	httpServer.WithContextInjectors(
		db.InjectDBInto(masterDB.GormDB()),
	)

	httpServer.Register(apis.RouterGroup_Root)
	return httpServer
}

func AutoMigrate() {
	err := migration.Magrate(masterDB)
	if err != nil {
		panic(err)
	}

}
