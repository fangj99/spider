package spider

import (
	"net"
	"sync"
)

//define const
const (
	TableSize    = 1024
	HasFoundSize = 100000
)

var hasFound = make(map[string]bool, HasFoundSize)
var mutex sync.Mutex

//KNode define
type KNode struct {
	ID   ID
	IP   net.IP
	Port int
}

//KTable define table
type KTable struct {
	Nodes []*KNode
	cap   int64

	//用于响应find_node请求
	Snodes []*KNode
}

//Put node to table
func (table *KTable) Put(node *KNode) {
	ids := string([]byte(node.ID))
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := hasFound[ids]; !ok {
		table.Nodes = append(table.Nodes, node)
		hasFound[ids] = true
	}
	if len(table.Snodes) < 8 {
		table.Snodes = append(table.Snodes, node)
	}
}

//Pop node
func (table *KTable) Pop() *KNode {
	if len(table.Nodes) > 0 {
		n := table.Nodes[0]
		table.Nodes = table.Nodes[1:]
		return n
	}
	return nil
}
