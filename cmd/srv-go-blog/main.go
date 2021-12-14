package main

import "github.com/tangx/srv-go-blog/cmd/srv-go-blog/global"

func main() {

	global.App.AddCommand("migrate", func(args ...string) {
		global.AutoMigrate()
	})

	global.App.Run(
		global.Server(),
	)
}
