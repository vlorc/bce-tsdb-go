package tsdb

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"net/url"
)

const (
	URI_DATAPOINT = "/v1/datapoint"
	URI_TAG       = "/v1/metric/%s/tag"
	URI_FIELD     = "/v1/metric/%s/field"
	URI_METRIC    = "/v1/metric"
)

type Client struct {
	*bce.BceClient
	host string
}

func NewClient(ak, sk, endpoint string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	credentials, err = auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(endpoint)
	if nil != err {
		return nil, err
	}

	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:    endpoint,
		Region:      bce.DEFAULT_REGION,
		UserAgent:   bce.DEFAULT_USER_AGENT,
		Credentials: credentials,
		SignOption:  defaultSignOptions,
		Retry:       bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	return &Client{BceClient: bce.NewBceClient(defaultConf, &auth.BceV1Signer{}), host: u.Host}, nil
}

func (c *Client) WriteDatapoint(data []Datapoint) error {
	return c.__post(URI_DATAPOINT, http.POST, &WriteDataPointArgs{DataPoints: data}, nil)
}

func (c *Client) __get(uri string, result interface{}, params ...string) (err error) {
	req := &bce.BceRequest{}
	res := &bce.BceResponse{}
	req.SetMethod(http.GET)
	req.SetUri(uri)
	for i, l := 0, len(params)/2; i < l; i++ {
		req.SetParam(params[i*2+0], params[i*2+1])
	}
	if err = c.SendRequest(req, res); nil != err {
		return
	}
	if res.IsFail() {
		err = res.ServiceError()
	} else {
		err = res.ParseJsonBody(result)
	}
	return
}

func (c *Client) __post(uri, method string, data, result interface{}, params ...string) (err error) {
	req := &bce.BceRequest{}
	req.SetMethod(method)
	req.SetUri(uri)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	for i, l := 0, len(params)/2; i < l; i++ {
		req.SetParam(params[i*2+0], params[i*2+1])
	}
	buf, err := json.Marshal(data)
	if nil != err {
		return
	}
	body, err := bce.NewBodyFromBytes(buf)
	if nil != err {
		return
	}
	req.SetBody(body)
	res := &bce.BceResponse{}
	if err = c.SendRequest(req, res); nil != err {
		return
	}
	if res.IsFail() {
		err = res.ServiceError()
	} else if nil != result {
		err = res.ParseJsonBody(result)
	}
	return
}

func (c *Client) ListMetric() ([]string, error) {
	list := &ListMetricsResult{}
	return list.Metrics, c.__get(URI_METRIC, list)
}

func (c *Client) ListFieldByMetric(metric string) (map[string]Field, error) {
	list := &ListFieldResult{}
	return list.Fields, c.__get(fmt.Sprintf(URI_FIELD, metric), list)
}

func (c *Client) ListTagByMetric(metric string) (map[string]TagValues, error) {
	list := &ListTagsResult{}
	return list.Tags, c.__get(fmt.Sprintf(URI_TAG, metric), list)
}

func (c *Client) ListDatapointByQuery(query Queries, disablePresampling ...bool) ([]QueryResult, error) {
	list := &ListDatapointResult{}
	return list.Results, c.__post(
		URI_DATAPOINT,
		http.PUT,
		&ListDatapointArgs{Queries: query, DisablePresampling: len(disablePresampling) > 0 && disablePresampling[0]},
		list, "query", "")
}

func (c *Client) ListRowBySql(statement string) (*RowResult, error) {
	list := &RowResult{}
	err := c.__get(URI_DATAPOINT, list, "sql", statement)
	if nil != err {
		list = nil
	}
	return list, err
}

func (c *Client) GeneratePresignedUrl(query Queries, expireSeconds int, endpoint ...string) (string, error) {
	buf, err := json.Marshal(&ListDatapointArgs{Queries: query})
	if nil != err {
		return "", err
	}
	req := &http.Request{}
	req.SetParam("query", string(buf))
	req.SetHeader("Host", c.host)
	req.SetUri(URI_DATAPOINT)
	req.SetMethod(http.GET)
	c.Signer.Sign(req, c.Config.Credentials, &auth.SignOptions{
		HeadersToSign: c.Config.SignOption.HeadersToSign,
		ExpireSeconds: expireSeconds,
	})

	val := url.Values{
		"query":         []string{req.Param("query")},
		"authorization": []string{req.Header(http.AUTHORIZATION)},
	}
	if len(endpoint) > 0 {
		return fmt.Sprintf("%s%s?%s", endpoint[0], URI_DATAPOINT, val.Encode()), nil
	}
	return fmt.Sprintf("%s%s?%s", c.Config.Endpoint, URI_DATAPOINT, val.Encode()), nil
}
