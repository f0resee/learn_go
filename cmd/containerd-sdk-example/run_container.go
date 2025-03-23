package main

import (
	"context"
	"fmt"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/oci"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func main() {
	client, err := containerd.New(
		"/run/containerd/containerd.sock",
		containerd.WithDefaultNamespace("default"),
	)
	if err != nil {
		log.Fatalf("init containerd client error: %s", err.Error())
	}
	defer client.Close()

	ctx := context.Background()

	image, err := client.GetImage(ctx, "docker.io/library/nginx:alpine")
	/* pull image with containerd sdk
	image, err := client.Pull(
		ctx,
		"docker.io/library/nginx:alpine",
		containerd.WithPullUnpack)
	*/
	if err != nil {
		log.Fatalf("failed to pull image: %s", err.Error())
	}

	hostBinaryPath := "/home/nansan/go/src/gitee/learn_go/cmd/containerd-sdk-example/run_container"
	containerBinaryPath := "/usr/bin/stress"

	mounts := []specs.Mount{
		{
			Type:        "bind",
			Source:      hostBinaryPath,
			Destination: containerBinaryPath,
			Options: []string{
				"bind",
			},
		},
	}

	container, err := client.NewContainer(ctx,
		"my-nginx",
		containerd.WithImage(image),
		containerd.WithNewSnapshot("my-nginx", image),
		containerd.WithNewSpec(
			oci.WithUsername("root"),
			oci.WithImageConfig(image),
			oci.WithMounts(mounts),
			oci.WithCPUs("0,4"),
			oci.WithEnv([]string{"MY_ENV_VAR=test"}),
		),
	)

	if err != nil {
		log.Fatalf("failed to create container: %s", err.Error())
	}

	task, err := container.NewTask(ctx, cio.NullIO)
	if err != nil {
		log.Fatalf("failed to create task: %s", err.Error())
	}

	if err := task.Start(ctx); err != nil {
		log.Fatalf("failed to start task: %s", err)
	}

	status, err := task.Wait(ctx)
	if err != nil {
		log.Fatalf("failed to wait for task: %s", err.Error())
	}

	exitStatus := <-status
	fmt.Printf("Containe exit with code: %d, status: %s\n", exitStatus.ExitCode(), exitStatus.Error())

}
