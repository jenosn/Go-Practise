package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Method     string
	StatusCode int
	Path       string
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w, Req: req, Method: req.Method, Path: req.URL.Path,
	}
}

//指定获取POST请求中的参数
func (context *Context) PostFrom(key string) string {
	return context.Req.PostFormValue(key)
}

//指定获取GET请求中的参数
func (context *Context) queryFrom(key string) string {
	return context.Req.Form.Get(key)

}

//当前挡球状态赋值
func (context *Context) Status(code int) {
	context.StatusCode = code
}

func (context *Context) setHeader(key string, value string) {
	context.Writer.Header().Set(key, value)
}

//返回string
func (context *Context) String(code int, format string, value ...interface{}) {
	context.setHeader("Content-Type", "text/plain")
	context.Status(code)
	context.Writer.Write([]byte(fmt.Sprintf(format, value)))
}

//返回json
func (context *Context) Json(code int, obj interface{}) {
	context.setHeader("Content-Type", "application/json")
	context.Status(code)
	encoder := json.NewEncoder(context.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.Writer, err.Error(), 500)
	}

}

//返回字节
func (context *Context) Data(code int, data []byte) {
	context.Status(code);
	context.Writer.Write(data)
}

func (context *Context) HTML(code int, html string){
	context.setHeader("Content-Type", "text/json")
	context.Status(code)
	context.Writer.Write([]byte(html))
}
