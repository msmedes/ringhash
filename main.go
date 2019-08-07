package main

import (
	"fmt"
	"math/rand"

	"github.com/msmedes/ringhash/ring"
)

func main() {
	distribution()
	rebalance()
}

func distribution() {
	nodes := []string{"a", "b", "c"}
	config := &ring.Config{
		VirtualNodes: 300,
		LoadFactor:   2,
	}
	r := ring.NewRing(nodes, config)

	distributeKeys(r)
}

func rebalance() {
	nodes := []string{"a", "b", "c"}
	config := &ring.Config{
		VirtualNodes: 300,
		LoadFactor:   2,
	}
	r := ring.NewRing(nodes, config)

	fmt.Println("Distributing keys to 3 node ring")
	distributeKeys(r)
	fmt.Println("Adding new node d")
	r.Add("d")
	fmt.Println("redistributing")
	distributeKeys(r)

}

func distributeKeys(r *ring.Ring) {
	// the amount of requests to generate
	requestCount := 1000000
	distribution := make(map[string]int)

	// keys will be len 4 byte arrays, which should give us 255^4 combinations
	key := make([]byte, 4)

	for i := 0; i < requestCount; i++ {
		// generate a random key
		rand.Read(key)
		// make a request
		node, err := r.Get(string(key))
		if err != nil {
			fmt.Println(err)
			continue
		}
		r.Finished(node)
		distribution[node]++
	}
	for node, count := range distribution {
		fmt.Printf("node: %s count: %d\n", node, count)
	}
}
