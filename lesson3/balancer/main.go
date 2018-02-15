package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	balancer := new(RoundRobinBalancer)
	balancer.Init(3)
	fmt.Println(balancer.GiveStat())

	balancer.GiveNode()
	balancer.GiveNode()
}

// RoundRobinBalancer is a load balancer
type RoundRobinBalancer struct {
	sync.Mutex
	nodes []*Node
}

// Node is server to handle request
type Node struct {
	requestCount int
	nodeID       int
}

func (b *RoundRobinBalancer) Len() int {
	return len(b.nodes)
}

func (b *RoundRobinBalancer) Less(i, j int) bool {
	return b.nodes[i].requestCount < b.nodes[j].requestCount
}

func (b *RoundRobinBalancer) Swap(i, j int) {
	b.nodes[i], b.nodes[j] = b.nodes[j], b.nodes[i]
}

// Init make connection with count of server
func (b *RoundRobinBalancer) Init(serverCount int) {
	b.nodes = make([]*Node, 0, serverCount)
	for i := 0; i < serverCount; i++ {
		b.nodes = append(b.nodes, &Node{0, i})
	}
}

// GiveStat make connection with count of server
func (b *RoundRobinBalancer) GiveStat() []int {
	result := make(sort.IntSlice, 0, len(b.nodes))
	for _, node := range b.nodes {
		result = append(result, node.requestCount)
	}
	sort.Sort(sort.Reverse(result))
	return []int(result)
}

// GiveNode make request to server
func (b *RoundRobinBalancer) GiveNode() int {
	b.Lock()
	defer b.Unlock()
	sort.Sort(b)
	b.nodes[0].requestCount++

	return b.nodes[0].nodeID
}
