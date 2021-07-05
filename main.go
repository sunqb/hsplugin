package main

import "hsplugin/server/webserver"

func main() {


	// webServer启动
	webserver.Init().Route(&webserver.FileController{}).Listen(":7080")
}
