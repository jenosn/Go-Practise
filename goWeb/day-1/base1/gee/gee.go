package gee

import (
	"fmt"
	"net/http"
)

type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path= %s\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Handler[%s]=%s\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not Found Url=%s", r.URL.Path)
	}

}


/*func main() {
	engine := new(Engine)
	http.ListenAndServe("9999", engine)
}
*/