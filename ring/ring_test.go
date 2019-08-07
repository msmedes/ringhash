package ring

import (
	"fmt"
	"math/rand"
	"testing"
)

// func TestGet(t *testing.T) {
// 	nodes := []string{"a", "b", "c"}
// 	cnf := &Config{
// 		VirtualNodes: 0,
// 		LoadFactor:   1,
// 	}
// 	hashRing := NewRing(nodes, cnf)

// 	expectNodesABC(t, hashRing)
// }

// func expectNodesABC(t *testing.T, hashRing *Ring) {
// 	// Python hash_ring module test case
// 	expectNode(t, hashRing, "test", "a")
// 	expectNode(t, hashRing, "test", "a")
// 	expectNode(t, hashRing, "test1", "c")
// 	expectNode(t, hashRing, "test2", "c")
// 	expectNode(t, hashRing, "test3", "c")
// 	expectNode(t, hashRing, "test4", "c")
// 	expectNode(t, hashRing, "test5", "b")
// 	expectNode(t, hashRing, "aaaa", "c")
// 	expectNode(t, hashRing, "bbbb", "a")
// }

// func expectNode(t *testing.T, hashRing *Ring, key string, expected string) {
// 	node, err := hashRing.Get(key)
// 	if err != nil || node != expected {
// 		t.Error("GetNode(", key, ") expected", expected, "but got", node, err)
// 	}
// 	hashRing.Done(node)
// }

// func failNode(t *testing.T, hashRing *Ring, key string, expected string, expectedErr error) {
// 	node, err := hashRing.Get(key)
// 	if err != expectedErr || node != expected {
// 		t.Error("GetNode(", key, ") expected", expected, "but got", node, err)
// 	}
// }

// func TestHeavyLoad(t *testing.T) {
// 	nodes := []string{"a", "b", "c"}
// 	cnf := &Config{
// 		VirtualNodes: 0,
// 		LoadFactor:   1,
// 	}
// 	hashRing := NewRing(nodes, cnf)
// 	failNode(t, hashRing, "test", "a", nil)
// 	failNode(t, hashRing, "test", "b", nil)
// 	failNode(t, hashRing, "test", "c", nil)
// 	failNode(t, hashRing, "test", "", ERR_HEAVY_LOAD)

// }

// func TestDistribution(t *testing.T) {
// 	nodes := []string{"a", "b", "c"}
// 	cnf := &Config{
// 		VirtualNodes: 0,
// 		LoadFactor:   1,
// 	}
// 	hashRing := NewRing(nodes, cnf)
// 	failNode(t, hashRing, "test", "a", nil)
// 	failNode(t, hashRing, "test", "b", nil)
// 	failNode(t, hashRing, "test", "c", nil)

// }

func TestDistribution(t *testing.T) {
	nodes := []string{"a", "b", "c"}
	cnf := &Config{
		VirtualNodes: 300,
		LoadFactor:   2,
	}
	hashRing := NewRing(nodes, cnf)

	keyCount := 1000000
	distribution := make(map[string]int)
	key := make([]byte, 4)
	for i := 0; i < keyCount; i++ {
		rand.Read(key)
		node, err := hashRing.Get(string(key))
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		hashRing.Finished(node)
		distribution[node]++
	}
	for node, count := range distribution {
		fmt.Printf("node: %s, key count: %d\n", node, count)
	}

}
