package tsdb

import "testing"

var __url = "http://xxxxx.tsdb.iot.bj.baidubce.com"
var __ak = "ssssssssssssssssssssssssssssssss"
var __sk = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

func Test_WriteData(t *testing.T) {
	cli, err := NewClient(__ak, __sk, __url)
	if nil != err {
		t.Error("NewClient: ", err)
		return
	}

	err = cli.WriteDatapoint([]Datapoint{{
		Metric: "cpu_idle",
		Tags: Tags{
			"host": "server1",
			"rack": "rack1",
		},
		Value: 51,
	}})
	if nil != err {
		t.Error("WriteDatapoint: ", err)
		return
	}
}

func Test_ListMetric(t *testing.T) {
	cli, err := NewClient(__ak, __sk, __url)
	if nil != err {
		t.Error("NewClient: ", err)
		return
	}

	metric, err := cli.ListMetric()
	if nil != err {
		t.Error("ListMetric: ", err)
		return
	}
	t.Log("metric: ", metric)
}

func Test_ListFieldByMetric(t *testing.T) {
	cli, err := NewClient(__ak, __sk, __url)
	if nil != err {
		t.Error("NewClient: ", err)
		return
	}

	field, err := cli.ListFieldByMetric("cpu_idle")
	if nil != err {
		t.Error("ListFieldByMetric: ", err)
		return
	}
	t.Log("field: ", field)
}

func Test_ListTagByMetric(t *testing.T) {
	cli, err := NewClient(__ak, __sk, __url)
	if nil != err {
		t.Error("NewClient: ", err)
		return
	}

	tag, err := cli.ListTagByMetric("cpu_idle")
	if nil != err {
		t.Error("ListTagByMetric: ", err)
		return
	}
	t.Log("tag: ", tag)
}

func Test_ListDatapoint(t *testing.T) {
	cli, err := NewClient(__ak, __sk, __url)
	if nil != err {
		t.Error("NewClient: ", err)
		return
	}

	data, err := cli.ListDatapointByQuery(Queries{{
		Metric: "cpu_idle",
		Limit:  10,
		Filters: Filter{
			Start: "12 hour ago",
		},
	}})
	if nil != err {
		t.Error("ListDatapointByQuery: ", err)
		return
	}
	t.Log("datapoint: ", data)
}

func Test_GeneratePresignedUrl(t *testing.T) {
	cli, err := NewClient(__ak, __sk, __url)
	if nil != err {
		t.Error("NewClient: ", err)
		return
	}

	url, err := cli.GeneratePresignedUrl(Queries{{
		Metric: "cpu_idle",
		Limit:  1000,
		Filters: Filter{
			Start: "12 hour ago",
		},
	}},1800)
	if nil != err {
		t.Error("GeneratePresignedUrl: ", err)
		return
	}
	t.Log("url: ", url)
}
