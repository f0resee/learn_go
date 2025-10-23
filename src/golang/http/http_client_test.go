package http

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"testing"
)

func Test_port_Client(t *testing.T) {
	client := &http.Client{}
	res, err := client.Get("http://localhost:8080/about")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(body))
}

func Test_uds_Client(t *testing.T) {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/http.sock")
			},
		},
	}
	res, err := client.Get("http://localhost/about")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(body))
}

func TestGoHttpGet(t *testing.T) {
	res, err := http.Get("https://www.baidu.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("res: ", res)
	log.Println("status: ", res.Status)
	log.Println("status code: ", res.StatusCode)
	log.Println("proto: ", res.Proto)
	log.Println("major proto: ", res.ProtoMajor)
	log.Println("minor proto:", res.ProtoMinor)
	log.Print("header: ", res.Header)
	log.Println("content-length: ", res.ContentLength)
	log.Println("close: ", res.Close)
	log.Println("uncompressed: ", res.Uncompressed)
	log.Println("trailer: ", res.Trailer)
	log.Println("transfer encoding", res.TransferEncoding)
	log.Println("request: ", res.Request)
	log.Println("tls: ", res.TLS)
	for k, v := range res.Header {
		log.Printf("(k:v) %s : %s\n", k, v)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code:%d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

func TestGoHttpRequestAndDefaultClient(t *testing.T) {
	var resReader http.Response
	req, err := http.NewRequest(http.MethodGet, "https://www.baidu.com/robots.txt", resReader.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		log.Fatal(err1)
	}
	body, err2 := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Printf("%s", body)
}
