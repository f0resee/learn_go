package http

import (
	"net"
	"net/http"
	"testing"
)

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	case "/about":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("About page"))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
	}
})

func Test_port_Server(t *testing.T) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	server.ListenAndServe()
}

// curl --unix-socket /tmp/http.sock http://localhost/about
func Test_uds_Server(t *testing.T) {
	listener, err := net.Listen("unix", "/tmp/http.sock")
	if err != nil {
		t.Fatal(err)
	}
	server := &http.Server{
		Handler: handler,
	}
	server.Serve(listener)
}
