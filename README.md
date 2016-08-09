# DHT网络爬虫
[![Build Status](https://drone.io/github.com/btlike/spider/status.png)](https://drone.io/github.com/btlike/spider/latest)

高性能DHT网络爬虫，5美金的单核,768MB的[VPS](https://www.vultr.com/pricing/)上，每秒处理UDP请求超过12K，内存占用不超过100MB，每天抓取数千万去重infohash

## 主要特点

- 内存复用
- map去重
- id 均匀分散
- 动态调整find_node速率
- 限速

## 示例

参考[example](https://github.com/btlike/spider/blob/master/example)


### 安装
`go get github.com/btlike/spider`



## 参考项目

- [dhtspider](https://github.com/alanyang/dhtspider)
- [DhtCrawler](https://github.com/xiaojiong/DhtCrawler)


## 流量图
![附一张流量图](https://github.com/btlike/spider/blob/master/flow.jpg)


## 常见问题
终于运行起了爬虫，但运行没几分钟，各种linux问题出现了，最开始应该是ulimit问题，这个问题很好解决，参考[这个文章](http://www.stutostu.com/?p=1322)。然后会出现开始大量报出：`nf_conntrack: table full, dropping packet`。这个问题参考[这个文章](http://jaseywang.me/2012/08/16/%E8%A7%A3%E5%86%B3-nf_conntrack-table-full-dropping-packet-%E7%9A%84%E5%87%A0%E7%A7%8D%E6%80%9D%E8%B7%AF/)。原因就是，

```
nf_conntrack/ip_conntrack 跟 nat 有关，用来跟踪连接条目，它会使用一个哈希表来记录 established 的记录。nf_conntrack 在 2.6.15 被引入，而 ip_conntrack 在 2.6.22 被移除，如果该哈希表满了，就会出现：nf_conntrack: table full, dropping packet。
```

解决办法很简单，我们让某些端口的流量不要被记录即可。假如我们运行100个节点，而节点监听的端口是20000到20099，我们只需要执行以下命令即可。

```
iptables -A INPUT -m state --state UNTRACKED -j ACCEPT
iptables -t raw -A PREROUTING -p udp -m udp --dport 20000 -j NOTRACK
...... //从端口20000一直到20099，每个端口一行
iptables -t raw -A PREROUTING -p udp -m udp --dport 20099 -j NOTRACK
```
