package main

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	numOfGoRoutine int
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "stress",
		Short: "stress",
		Long:  "stress test",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("running with %d goroutine\n", numOfGoRoutine)
			var wg sync.WaitGroup
			for i := 0; i < numOfGoRoutine; i++ {
				wg.Add(1)
				go func() {
					j := 1
					for true {
						j++
						if j%10000000000 == 0 {
							time.Sleep(1 * time.Millisecond)
						}
					}
					wg.Done()
				}()
			}
			wg.Wait()
			return nil
		},
	}
	rootCmd.PersistentFlags().IntVar(&numOfGoRoutine, "goroutine", 1, "number of goroutine")
	err := rootCmd.Execute()
	log.Fatalf("%+v\n", err)
}
