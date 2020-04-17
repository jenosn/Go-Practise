package main

import (
	"awesomeProject/goWeb/day-1/base2/gee"
	"fmt"
	"net/http"
)

func main() {
	engine := gee.New();
	engine.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		fmt.Fprintf(w, "URL.Path=%s", r.URL.Path)
	})
	engine.Get("hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "route[%s]=%s", k, v)
		}
	})
	engine.Run(":9999")
}
