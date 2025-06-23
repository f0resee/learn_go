package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/json"
)

type RuntimeOptions struct {
	Address   string
	Namespace string
}

func (r *RuntimeOptions) Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("runtime", pflag.ExitOnError)
	fs.StringVarP(&r.Address, "address", "a", "/var/run/containerd/containerd.sock", "containerd sock address")
	fs.StringVarP(&r.Namespace, "namespace", "n", "default", "containerd namespace")
	return fs
}

var (
	runtimeOptions RuntimeOptions

	env string

	timeout int64
)

func PrintFlags(location string) {
	fmt.Printf("in %v, env: %s, runtime option: %s, timeout:%d\n", location, env, ToS(runtimeOptions), timeout)
}

func main() {
	// 1. new rootCmd
	rootCmd := &cobra.Command{
		Use: "ctr [-a path] [-n namespace] c ls",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			PrintFlags("PersistentPreRunE")
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			PrintFlags("PreRunE")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			PrintFlags("RunE")
			return nil
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			PrintFlags("PostRunE")
			return nil

		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			PrintFlags("PersistentPostRunE")
			return nil
		},
	}

	// 2. add a global flag name env to cmd
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "prod", "env name")
	PrintFlags("after StringVarP")

	// 3. add flags through a struct
	fs := rootCmd.Flags()
	fs.AddFlagSet(runtimeOptions.Flags())

	// 4. add a flag through FlagSet
	afs := pflag.NewFlagSet("time", pflag.ExitOnError)
	afs.Int64VarP(&timeout, "timeout", "t", 0, "operation timeout in seconds")
	fs.AddFlagSet(afs)

	// 5. add sub cmd
	rootCmd.AddCommand(&cobra.Command{
		Use: "container",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("in RunE: %+v\n", args)
			return nil
		},
	})

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Execite error: %s", err.Error())
	}
	PrintFlags("after Execute")
}

func ToS(value interface{}) string {
	data, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("marshal json err: %s", err.Error())
	}
	return string(data)
}
