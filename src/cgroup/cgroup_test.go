package cgroup

import (
	"os"
	"testing"
	"time"

	"github.com/containerd/cgroups"
	cgroupsv2 "github.com/containerd/cgroups/v2"
	_ "github.com/opencontainers/runtime-spec/specs-go"
)

func TestCgoupV2(t *testing.T) {
	var cgroupV2 bool
	if cgroups.Mode() == cgroups.Unified {
		cgroupV2 = true
	}
	t.Logf("cgroupV2 = %v", cgroupV2)
}

func TestCgroupv2(t *testing.T) {
	res := cgroupsv2.Resources{
		CPU: &cgroupsv2.CPU{},
	}
	// dummy PID of -1 is used for creating a "general slice" to be used as a parent cgroup.
	// see https://github.com/containerd/cgroups/blob/1df78138f1e1e6ee593db155c6b369466f577651/v2/manager.go#L732-L735
	m, err := cgroupsv2.NewSystemd("/", "my-cgroup-abc.slice", -1, &res)
	if err != nil {
		t.Logf("err = %s", err.Error())
	} else {
		pid := os.Getpid()
		err := m.AddProc(uint64(pid))
		time.Sleep(100 * time.Second)
		m.Delete()
	}

}
