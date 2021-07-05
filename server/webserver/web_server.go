package webserver

import "github.com/gin-gonic/gin"

// 基于DDD思想设计的WebServer结构体

// WebServer web服务器结构体
type WebServer struct {
	router      *gin.Engine
	routerGroup *gin.RouterGroup
}

// Init 初始化WebServer
func Init() *WebServer {
	webServer := &WebServer{router: gin.New()}
	return webServer
}

// Listen 在:xxx启动服务。default port is ":8080"
func (webServer *WebServer) Listen(port string) {
	if port != "" {
		webServer.router.Run(port)
	} else {
		webServer.router.Run()
	}
}

// Route 路由挂载。controller必须实现IController方法，在方法内实现具体路由
func (webServer *WebServer) Route(controllers ...IController) *WebServer {
	for _, c := range controllers {
		c.Router(webServer)
	}
	return webServer
}

// GroupRouter 路由分组
func (webServer *WebServer) GroupRouter(group string, controllers ...IController) *WebServer {
	webServer.routerGroup = webServer.router.Group(group)
	for _, c := range controllers {
		c.Router(webServer)
	}
	return webServer
}
