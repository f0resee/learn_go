package main //package

import (
	"fmt"
	"os"
) //dependency

// go version
// go run hello_world.go

//entry:
//1. must be main package: package main
//2. must be function main: func main
//3. the file name may not be main.go

//main function cannot return any value
//return some state by using os.Exit

// main doesn't have parameter
func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello, world", os.Args[1])
	}
	fmt.Println("Hello, world!")
	os.Exit(0)
}
