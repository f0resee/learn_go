package gocontext

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		fmt.Println("dong. true")
		return true
	default:
		fmt.Println("not done. false")
		return false
	}
}

func TestContext(t *testing.T)  {
	ctx , cancel := context.WithCancel(context.Background())
	for i:=0;i<5;i++ {
		go func(i int,ctx context.Context) {
			for  {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond*250)
				t.Log(i,"WORKING")
			}
			t.Log(i," canceled")
		}(i,ctx)
	}
	fmt.Println("try to cancel")

	time.Sleep(time.Second*1)
	cancel()
}

func TestGoContext(t *testing.T)  {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// monitor
		go func(ctx context.Context) {
			for range time.Tick(time.Second) {
				select {
				case <-ctx.Done():
					fmt.Println("req is outgoing")
					return
				default:
					fmt.Println("req is processing")
				}
			}
		}(r.Context())

		// assume req processing takes 3s
		time.Sleep(3 * time.Second)
		w.Write([]byte("hello"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}