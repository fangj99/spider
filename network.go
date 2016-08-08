package spider

import (
	"net"
	"time"

	"github.com/juju/ratelimit"
)

//RateLimit limit speed
var RateLimit int64 = 100

//Network define
type Network struct {
	Dht       *DhtNode
	Conn      *net.UDPConn
	RateLimit *ratelimit.Bucket
}

//NewNetwork create network
func NewNetwork(dhtNode *DhtNode, address string) *Network {
	nw := new(Network)
	nw.Dht = dhtNode
	nw.RateLimit = ratelimit.NewBucketWithRate(float64(RateLimit), RateLimit) //默认限速：每个节点每秒最多处理100个请求
	nw.Init(address)
	return nw
}

//Init it
func (nw *Network) Init(address string) {
	addr := new(net.UDPAddr)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		panic(err)
	}
	nw.Conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	laddr := nw.Conn.LocalAddr().(*net.UDPAddr)
	nw.Dht.node.IP = laddr.IP
	nw.Dht.node.Port = laddr.Port
}

//Listening on
func (nw *Network) Listening() {
	val := make(map[string]interface{})
	buf := make([]byte, 1024)
	for {
		time.Sleep(nw.RateLimit.Take(1))
		n, raddr, err := nw.Conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		nw.Dht.krpc.Decode(buf[:n], val, raddr)
	}
}

//Send data
func (nw *Network) Send(m []byte, addr *net.UDPAddr) error {
	_, err := nw.Conn.WriteToUDP(m, addr)
	if err != nil {
		logger(err)
	}
	return err
}
