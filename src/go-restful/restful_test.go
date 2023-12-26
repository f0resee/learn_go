package go_restful

import (
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/emicklei/go-restful/v3"
)

func TestGoRestfullHello(t *testing.T) {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
