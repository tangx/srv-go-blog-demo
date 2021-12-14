package global

import (
	"github.com/go-jarvis/confhttp"
	"github.com/go-jarvis/jarvis"
	"github.com/go-jarvis/jarvis/pkg/appctx"
	"github.com/tangx/srv-go-blog/cmd/srv-go-blog/apis"
)

var (
	httpServer = &confhttp.Server{}
	//masterDB   = &drivers.MysqlDriver{}
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
		//MasterDB   *drivers.MysqlDriver
	}{
		HttpServer: httpServer,
		//MasterDB:   masterDB,
	}

	_ = App.Conf(config)
}

func Server() *confhttp.Server {

	httpServer.Register(apis.RouterGroup_Root)
	return httpServer
}
