package main

import (
    "Go-Practise/goWeb/day-2/gee"
    "net/http"
)

func main() {
    engine := gee.NewEngine()
    engine.Get("/", func(ctx *gee.Context) {
        ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
    })
    engine.Run(":9999")

}
