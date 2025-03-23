package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	runtimev1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

var (
	sockAddr string

	conn           *grpc.ClientConn
	runtimeClient  runtimev1.RuntimeServiceClient
	criImageClient runtimev1.ImageServiceClient
)

func main() {

	sockAddr = "unix:///run/containerd/containerd.sock"

	log.Printf("critool trys to dial containerd sock addr: %s", sockAddr)
	sockConn, err := grpc.Dial(sockAddr, grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		log.Fatalf("fail to dial containerd sock: %s, err: %+v", sockAddr, err)
	}
	conn = sockConn
	runtimeClient = runtimev1.NewRuntimeServiceClient(conn)
	criImageClient = runtimev1.NewImageServiceClient(conn)

	ctx := context.Background()
	sandboxConfig := &runtimev1.PodSandboxConfig{
		Metadata: &runtimev1.PodSandboxMetadata{
			Name:      "pod-my11",
			Namespace: "default",
		},
		Hostname:     "pod-my11",
		LogDirectory: "/home/nansan/go/src/gitee/learn_go/cmd/cri-api-example",
		DnsConfig:    nil,
		Labels:       map[string]string{"app": "my-app"},
		Annotations:  nil,
		Linux: &runtimev1.LinuxPodSandboxConfig{
			CgroupParent: "kubepods-burstable.slice",
		},
	}
	sandbox, err := runtimeClient.RunPodSandbox(ctx, &runtimev1.RunPodSandboxRequest{
		Config: sandboxConfig,
	})
	if err != nil {
		log.Fatalf("fail to create pod sandbox: %s", err.Error())
	}
	log.Printf("podsandbox created: %s", sandbox.PodSandboxId)

	containerConfig := &runtimev1.ContainerConfig{
		Metadata: &runtimev1.ContainerMetadata{
			Name: "my-container11",
		},
		Image: &runtimev1.ImageSpec{
			Image: "docker.io/library/busybox:latest",
		},
		Command: []string{"sleep", "3600"},
		Tty:     false,
		LogPath: "/home/nansan/go/src/gitee/learn_go/cmd/cri-api-example",
	}

	container, err := runtimeClient.CreateContainer(ctx, &runtimev1.CreateContainerRequest{
		PodSandboxId:  sandbox.PodSandboxId,
		Config:        containerConfig,
		SandboxConfig: sandboxConfig,
	})
	if err != nil {
		log.Fatalf("fail to create container: %s", err.Error())
	}
	fmt.Printf("container created: %s\n", container.ContainerId)

	if _, err = runtimeClient.StartContainer(ctx, &runtimev1.StartContainerRequest{
		ContainerId: container.ContainerId,
	}); err != nil {
		log.Fatalf("fail to start container: %s", err.Error())
	}
	fmt.Println("container start successfully")

	status, err := runtimeClient.ContainerStats(ctx, &runtimev1.ContainerStatsRequest{
		ContainerId: container.ContainerId,
	})
	if err != nil {
		log.Fatalf("container status error: %s", err.Error())
	}
	log.Printf("status", status)
	time.Sleep(1 * time.Hour)
}
