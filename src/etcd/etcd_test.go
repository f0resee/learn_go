package etcd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func Test_Etcd(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Errorf("new etcd client error: %s", err.Error())
	}
	defer client.Close()

	ctx := context.Background()
	lease, err := client.Grant(ctx, 10)
	if err != nil {
		t.Errorf("etcd client lease err: %s", err.Error())
	}

	timeCtx, _ := context.WithTimeout(ctx, 5*time.Second)
	resp, err := client.Put(timeCtx, "/c", "c", clientv3.WithLease(lease.ID))
	if err != nil {
		t.Errorf("etcd put error: %s", err.Error())
	}
	t.Logf("etcd put resp: %+v", resp)

	resp1, err := client.Get(ctx, "/c")
	if err != nil {
		t.Errorf("etcd get error: %s", err.Error())
	}
	for _, kv := range resp1.Kvs {
		fmt.Printf("key: %s,value: %s\n", kv.Key, kv.Value)
	}
	t.Logf("after sleep")
	time.Sleep(12 * time.Second)
	resp1, err = client.Get(ctx, "/c")
	if err != nil {
		t.Errorf("get etcd error: %s", err.Error())
	}
	for _, kv := range resp1.Kvs {
		fmt.Printf("key: %s,value: %s\n", kv.Key, kv.Value)
	}
}
