package consul

import (
	"fmt"
	"github.com/Kidsunbo/kie_toolbox_go/logs"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestClient(t *testing.T) {
	//var lastIndex uint64
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500" //consul server

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed, err:", err)
		return
	}

	service, _, _ := client.Catalog().Service("serverNode", "", nil)
	if err != nil {
		logs.Info("err = %s", err.Error())
	} else {
		logs.Info("service = %v", service)
	}

	services, _, err := client.Health().Service("serverNode", "v1000", true, nil)
	/*
		services, metainfo, err := client.Health().Service("serverNode", "v1000", true, &api.QueryOptions{
			WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
		})
		if err != nil {
			logrus.Warnf("error retrieving instances from Consul: %s", err.Error())
		}
		lastIndex = metainfo.LastIndex

		addrs := map[string]struct{}{}
	*/
	for _, service := range services {
		fmt.Println("service.Service.Address:", service.Service.Address, "service.Service.Port:", service.Service.Port)
		//addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = struct{}{}
	}
}
