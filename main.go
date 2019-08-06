package main

import (
	"fmt"

	"github.com/msmedes/ringhash/ring"
)

func main() {
	nodes := []string{"leeloo", "ruby", "mike"}
	config := &ring.Config{VirtualNodes: 200, LoadFactor: 2.0}
	r := ring.NewRing(nodes, config)
	fmt.Println(len(r.NodeMap()))
	fmt.Println(r.Get("leeloo"))
	r.Finished("leeloo")
	r.Remove("ruby")
	r.Get("ruby")
	fmt.Println(len(r.NodeMap()))
}
