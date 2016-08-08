# DHT网络爬虫

高性能DHT网络爬虫，5美金的单核,768MB的[VPS](https://www.vultr.com/pricing/)上，每秒处理UDP请求超过12K，内存占用不超过100MB，每天抓取数千万去重infohash

## 主要特点

- 内存复用
- map去重
- id 均匀分散
- 动态调整find_node速率
- 限速

## 示例

参考[example](https://github.com/btlike/spider/blob/master/example)



## 参考项目

- [dhtspider](https://github.com/alanyang/dhtspider)
- [DhtCrawler](https://github.com/xiaojiong/DhtCrawler)



![附一张流量图](https://github.com/btlike/spider/blob/master/flow.jpg)