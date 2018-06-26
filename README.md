# 百度TSDB GO SDK
时序数据库基本操作，API文档参考[官方](https://cloud.baidu.com/doc/TSDB/API.html#.E6.95.B0.E6.8D.AEAPI.E6.8E.A5.E5.8F.A3.E8.AF.B4.E6.98.8E)
## 安装
```shell
go get github.com\vlorc\bce-tsdb-go
```

## 功能
+ WriteDatapoint: 写入data point
+ ListMetric: 获取metric列表
+ ListFieldByMetric: 获取field列表
+ ListTagByMetric: 获取tag列表
+ ListDatapointByQuery: 查询data point
+ GeneratePresignedUrl: 生成查询URL

##　例子
1. 创建客户端

```go
import "github.com\vlorc\bce-tsdb-go"

func main() {
	// 创建STS服务的Client对象，Endpoint使用默认值
	AK, SK := <your-access-key-id>, <your-secret-access-key>
	//指明使用HTTPS协议
	ENDPOINT := "https://xxxxx.tsdb.iot.bj.baidubce.com"
	cli, err := tsdb.NewClient(AK, SK,ENDPOINT)
}
```

2. 写入数据

```go
err = cli.WriteDatapoint([]Datapoint{{
	Metric: "cpu_idle",
	Tags: Tags{
		"host": "server1",
		"rack": "rack1",
	},
	Value: 51,
}})
```