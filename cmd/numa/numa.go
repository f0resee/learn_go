package main

import (
	"log"
	"net/http"
	"time"

	"github.com/google/cadvisor/cache/memory"
	"github.com/google/cadvisor/container"
	"github.com/google/cadvisor/manager"
	"github.com/google/cadvisor/utils/sysfs"
	"k8s.io/apimachinery/pkg/util/json"
)

func main() {
	sysFs := sysfs.NewRealSysFs()
	duration := 15 * time.Second
	allowDynamic := true
	houseKeepConfig := manager.HousekeepingConfig{
		Interval:     &duration,
		AllowDynamic: &allowDynamic,
	}
	includeMetrics := container.MetricSet{}

	cgroupRoots := []string{"/sys/fs/cgroup"}

	m, err := manager.New(memory.New(2*time.Minute, nil), sysFs, houseKeepConfig, includeMetrics, http.DefaultClient, cgroupRoots, nil, "", time.Duration(0))
	if err != nil {
		log.Fatalf("new manager err: %+v", err)
	}
	machineInfo, err := m.GetMachineInfo()
	if err != nil {
		log.Fatalf("get machine info error: %+v", err)
	}
	data, _ := json.Marshal(machineInfo)
	log.Printf("get machine info: %s", string(data))
}
