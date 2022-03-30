package main

import (
	"fmt"

	"github.com/karlssonerik/gobworkers"
)

func main() {
	workers := gobworkers.New(5)
	for i := 0; i < 40; i++ {
		j := i
		workers.AddWork(func() { fmt.Println(j) })
	}

	workers.WaitForWorkToBeDone()
}
