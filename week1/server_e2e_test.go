//go:build e2e

package week1

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	h := &HTTPServer{}

	h.AddRoute(http.MethodGet, "/user", func(ctx Context) {
		fmt.Println("处理完")
	})

	h.Get("/user", func(ctx Context) {
		fmt.Println("我是get")
	})

	//这个handler就是我们和http包的结合点
	//用法一
	//http.ListenAndServe(":8088", h)
	//http.ListenAndServeTLS(":8088", "", "", h)

	//用法二
	h.Start(":8082")
}
