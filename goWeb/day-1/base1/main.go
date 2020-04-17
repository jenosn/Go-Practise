package main

import (
	"awesomeProject/goWeb/day-1/base1/gee"
	"net/http"
)

func main() {
	engine := new(gee.Engine)
	http.ListenAndServe(":9999", engine)
}
