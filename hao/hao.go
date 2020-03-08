package hao

import (
	"net/http"
)

// 定义hao使用的请求处理程序
type HandlerFunc func(ctx *Context)

// 实现 serverHttp接口
type Engine struct {
	router *router
}

//  这是一个Engine 的构造器
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 添加路由
func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET 定义处理get请求发方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 定义处理POST请求的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 定义启动HTTP服务器的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
