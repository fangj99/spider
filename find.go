package spider

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/zeebo/bencode"
)

//BOOTSTRAP define
var BOOTSTRAP = []string{
	"67.215.246.10:6881",
	"212.129.33.50:6881",
	"82.221.103.244:6881"}

var findDelayTime = time.Millisecond * 50

//FindNode find node
func (dhtNode *DhtNode) FindNode(v map[string]interface{}, args map[string]string, node *KNode) {
	countFindRequest++

	var id ID
	if node.ID != nil {
		id = node.ID.Neighbor(dhtNode.node.ID)
	} else {
		id = dhtNode.node.ID
	}

	v["t"] = fmt.Sprintf("%d", rand.Intn(100))
	v["y"] = "q"
	v["q"] = "find_node"

	args["id"] = string(id)
	args["target"] = string(GenerateID())
	v["a"] = args
	data, err := bencode.EncodeBytes(v)
	if err != nil {
		logger(err)
		return
	}

	raddr := new(net.UDPAddr)
	raddr.IP = node.IP
	raddr.Port = node.Port
	err = dhtNode.network.Send(data, raddr)
	if err != nil {
		logger(err)
		return
	}
}

//NodeFinder node finder
func (dhtNode *DhtNode) NodeFinder() {
	if len(dhtNode.table.Nodes) == 0 {
		val := make(map[string]interface{})
		args := make(map[string]string)
		for _, host := range BOOTSTRAP {
			raddr, err := net.ResolveUDPAddr("udp", host)
			if err != nil {
				logger("Resolve DNS error, %s\n", err)
				return
			}
			node := new(KNode)
			node.Port = raddr.Port
			node.IP = raddr.IP
			node.ID = nil
			dhtNode.FindNode(val, args, node)
		}
	}
	val := make(map[string]interface{})
	args := make(map[string]string)
	for {
		node := dhtNode.table.Pop()
		if node != nil {
			dhtNode.FindNode(val, args, node)
			time.Sleep(findDelayTime)
			continue
		}
		time.Sleep(time.Second)
	}
}
