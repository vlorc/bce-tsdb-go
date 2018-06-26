package tsdb

type ValueType string

const (
	Long   ValueType = "Long"
	Double           = "Double"
	String           = "String"
	Bytes            = "Bytes"
	Number           = "Number"
)

type Tags map[string]string

type Datapoint struct {
	Metric    string          `json:"metric"`
	Field     string          `json:"field,omitempty"`
	Tags      Tags            `json:"tags"`
	Type      ValueType       `json:"type,omitempty"`
	Timestamp int64           `json:"timestamp,omitempty"`
	Value     interface{}     `json:"value,omitempty"`
	Values    [][]interface{} `json:"values,omitempty"`
}

type WriteDataPointArgs struct {
	DataPoints []Datapoint `json:"datapoints"`
}

type ListMetricsResult struct {
	Metrics []string `json:"metrics"`
}

type Field struct {
	Type ValueType `json:"fields"`
}

type ListFieldResult struct {
	Fields map[string]Field `json:"fields"`
}

type TagValues []string

type ListTagsResult struct {
	Tags map[string]TagValues `json:"tags"`
}

type Query struct {
	Metric      string       `json:"metric"`
	Field       string       `json:"field,omitempty"`
	Fields      []string     `json:"fields,omitempty"`
	Tags        []string     `json:"tags,omitempty"`
	Filters     Filter       `json:"filters"`
	GroupBy     []GroupBy    `json:"groupBy,omitempty"`
	Limit       int          `json:"limit,omitempty"`
	Aggregators []Aggregator `json:"aggregators,omitempty"`
	Order       string       `json:"order,omitempty"`
	Fill        *Fill        `json:"fill,omitempty"`
	Fills       []Fill       `json:"fills,omitempty"`
	Marker      string       `json:"marker,omitempty"`
}

type Queries []Query

type ListDatapointArgs struct {
	Queries            Queries `json:"queries"`
	DisablePresampling bool    `json:"disablePresampling,omitempty"`
}

type ListDatapointResult struct {
	Results []QueryResult `json:"results"`
}

type QueryResult struct {
	Metric            string   `json:"metric"`
	Field             string   `json:"field"`
	Fields            []string `json:"fields"`
	Tags              []string `json:"tags"`
	RawCount          int      `json:"rawCount"`
	Groups            []Group  `json:"groups"`
	Truncated         bool     `json:"truncated"`
	NextMarker        string   `json:"nextMarker"`
	PresamplingRuleId string   `json:"presamplingRuleId"`
}

type Group struct {
	GroupInfos []GroupInfo     `json:"groupInfos"`
	Values     [][]interface{} `json:"values"`
}

type GroupInfo struct {
	Name string               `json:"name"`
	Tags map[string]TagValues `json:"tags"`
}

type Fill struct {
	Name             string `json:"type"`
	Interval         string `json:"interval"`
	MaxWriteInterval string `json:"maxWriteInterval,omitempty"`
}

type Aggregator struct {
	Name       string  `json:"name"`
	Sampling   string  `json:"sampling,omitempty"`
	Percentile float64 `json:"percentile,omitempty"`
	Divisor    float64 `json:"divisor,omitempty"`
	Factor     float64 `json:"factor,omitempty"`
	TimeUnit   string  `json:"timeUnit,omitempty"`
}

type GroupBy struct {
	Name string   `json:"name"`
	Tags []string `json:"tags,omitempty"`
}

type Filter struct {
	Start  interface{}          `json:"start"`
	End    interface{}          `json:"end,omitempty"`
	Tags   map[string]TagValues `json:"tags,omitempty"`
	Value  string               `json:"value,omitempty"`
	Fields []FieldFilter        `json:"fields,omitempty"`
	Or     []Filter             `json:"or,omitempty"`
}

type FieldFilter struct {
	Field string `json:"start"`
	Value string `json:"end"`
}

type TagFilter struct {
	Tag   string   `json:"tag"`
	In    []string `json:"in"`
	NotIn []string `json:"notIn"`
	Like  string   `json:"like"`
}
