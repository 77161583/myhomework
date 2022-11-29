package week1

import (
	"net"
	"net/http"
)

type handleFunc func(ctx Context)

// 确保一定实现了server接口
var _ Server = &HTTPServer{}

type Server interface {
	http.Handler
	Start(add string) error

	// AddRoute 增加路由注册功能
	// handleFunc 业务逻辑
	AddRoute(method string, path string, handleFunc handleFunc)
	//AddRoute1(method string, path string, handleFunc ...handleFunc)
}

type HTTPServer struct {
}

// ServerHTTP 处理请求的入口
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//框架代码就在这里
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {
	//接下来就是查找路由，并且执行命中路由的逻辑

}

func (h *HTTPServer) AddRoute(method string, path string, handleFunc handleFunc) {
	//panic("还没写")
}

func (h *HTTPServer) Get(path string, handleFunc handleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

func (h *HTTPServer) Post(path string, handleFunc handleFunc) {
	h.AddRoute(http.MethodPost, path, handleFunc)
}

//func (h *HTTPServer) AddRoute1(method string, path string, handleFunc ...handleFunc) {
//	panic("还没写1")
//}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	//在这里可以让用户注册所谓的after 和start 回调
	//比如说往 admin 注册一下自己的实例
	//在这里执行一些业务所需要的前置条件
	return http.Serve(l, h)
}

func (h *HTTPServer) Start1(addr string) error {
	return http.ListenAndServeTLS(addr, "", "", h)
}
