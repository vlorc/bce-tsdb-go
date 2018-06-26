# 百度TSDB GO SDK
# [Bce-tsdb-go](https://github.com/vlorc/bce-tsdb-go)

[![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![codebeat badge](https://codebeat.co/badges/c41b426c-4121-4dc8-99c2-f1b60574be64)](https://codebeat.co/projects/github-com-vlorc-bce-tsdb-go-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/vlorc/bce-tsdb-go)](https://goreportcard.com/report/github.com/vlorc/bce-tsdb-go)
[![GoDoc](https://godoc.org/github.com/vlorc/bce-tsdb-go?status.svg)](https://godoc.org/github.com/vlorc/bce-tsdb-go)
[![Build Status](https://travis-ci.org/vlorc/bce-tsdb-go.svg?branch=master)](https://travis-ci.org/vlorc/bce-tsdb-go?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/vlorc/bce-tsdb-go/badge.svg?branch=master)](https://coveralls.io/github/vlorc/bce-tsdb-go?branch=master)

百度时序数据库基本操作，API文档参考[官方](https://cloud.baidu.com/doc/TSDB/API.html#.E6.95.B0.E6.8D.AEAPI.E6.8E.A5.E5.8F.A3.E8.AF.B4.E6.98.8E)

## 安装
```shell
go get github.com\vlorc\bce-tsdb-go
```

## 许可证
这个项目是在Apache许可证下进行的。请参阅完整许可证文本的许可证文件。

## 功能
+ WriteDatapoint:  写入data point
+ ListMetric:  获取metric列表
+ ListFieldByMetric:  获取field列表
+ ListTagByMetric:  获取tag列表
+ ListDatapointByQuery:  查询data point
+ GeneratePresignedUrl:  生成查询URL

## 例子
1. 创建客户端

```go
import "github.com\vlorc\bce-tsdb-go"

func main() {
	// 创建TSDB服务的Client对象
	AK, SK := <your-access-key-id>, <your-secret-access-key>
	// 指明使用HTTPS协议
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