package webserver

// 路由控制

// IController 路由接口
type IController interface {
	Router(server *WebServer)
}
