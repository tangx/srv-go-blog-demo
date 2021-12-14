package main

import "github.com/tangx/srv-go-blog/cmd/srv-go-blog/global"

func main() {

	global.App.Run(
		global.Server(),
	)
}
