package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	runtimev1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

var (
	sockAddr string

	conn           *grpc.ClientConn
	runtimeClient  runtimev1.RuntimeServiceClient
	criImageClient runtimev1.ImageServiceClient
)

func newImageCommand() *cobra.Command {
	// 1. create image command
	imageCmd := &cobra.Command{
		Use:   "image",
		Short: "image is used to manage images",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	// 2. add image sub command: ls
	imageCmd.AddCommand(&cobra.Command{
		Use:   "ls",
		Short: "List images",
		RunE: func(cmd *cobra.Command, args []string) error {
			listImageResp, err := criImageClient.ListImages(context.TODO(), &runtimev1.ListImagesRequest{})
			if err != nil {
				log.Fatalf("failed to list images: %+v", err)
				return err
			}
			log.Printf("images: %+v", ToString(listImageResp.Images))
			return nil
		},
	})

	// 3. add image sub command: pull
	imageCmd.AddCommand(&cobra.Command{
		Use:   "pull",
		Short: "Pull image",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("must specify image name")
			}
			image := args[0]

			log.Printf("trying to pull image: %s", image)
			pulledImage, err := criImageClient.PullImage(context.TODO(), &runtimev1.PullImageRequest{
				Image: &runtimev1.ImageSpec{
					Image: image,
				},
			})
			if err != nil {
				log.Fatalf("failed to pull image: %s, err:%s", image, err)
				return err
			}
			log.Printf("pull image: %s success, pulledImage: %+v", imageCmd, pulledImage)
			return nil
		},
	})
	return imageCmd
}

func main() {
	// 1. rootCmd support --help
	rootCmd := &cobra.Command{
		Use:   "critool",
		Short: "critool is a CRI tool",
		Long:  "critool is a CRI tool, used to test CRI API",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("critool trys to dial containerd sock addr: %s", sockAddr)
			sockConn, err := grpc.Dial(sockAddr, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("fail to dial containerd sock: %s, err: %+v", sockAddr, err)
				return err
			}
			conn = sockConn
			runtimeClient = runtimev1.NewRuntimeServiceClient(conn)
			criImageClient = runtimev1.NewImageServiceClient(conn)
			return nil
		},
	}
	rootCmd.AddCommand(newImageCommand())
	rootCmd.PersistentFlags().StringVar(&sockAddr, "address", "unix:///run/containerd/containerd.sock", "containerd sock address")
	err := rootCmd.Execute()
	log.Fatalf("%v\n", err)
}
func ToString(o interface{}) string {
	data, _ := json.Marshal(o)
	return string(data)
}
