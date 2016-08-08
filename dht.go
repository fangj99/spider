package spider

import (
	"fmt"
)

//DhtNode define
type DhtNode struct {
	node    *KNode
	table   *KTable
	network *Network
	krpc    *KRPC
	outChan chan AnnounceData
}

//NewDhtNode create node
func NewDhtNode(id *ID, outHashIDChan chan AnnounceData, address string) *DhtNode {
	node := new(KNode)
	node.ID = *id
	dht := new(DhtNode)
	dht.outChan = outHashIDChan
	dht.node = node
	dht.table = new(KTable)
	dht.network = NewNetwork(dht, address)
	dht.krpc = NewKRPC(dht)
	return dht
}

//Run spider
func (dht *DhtNode) Run() {
	go func() { dht.network.Listening() }()
	go func() { dht.NodeFinder() }()
	logger(fmt.Sprintf("爬虫节点开始运行,最大处理速度每秒:%d,监听地址:%s", RateLimit, dht.network.Conn.LocalAddr().String()))
}
