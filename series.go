/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import "fmt"

// DataPoint is a tuple of [UNIX timestamp, value]. This has to use floats
// because the value could be non-integer.
type DataPoint [2]float64

// Metric represents a collection of data points that we might send or receive
// on one single metric line.
type Metric struct {
	Metric string      `json:"metric,omitempty"`
	Points []DataPoint `json:"points,omitempty"`
	Type   string      `json:"type,omitempty"`
	Host   string      `json:"host,omitempty"`
	Tags   []string    `json:"tags,omitempty"`
}

// reqPostSeries from /api/v1/series
type reqPostSeries struct {
	Series []Metric `json:"series"`
}

// PostSeries takes as input a slice of metrics and then posts them up to the
// server for posting data.
func (self *Client) PostMetrics(series []Metric) error {
	return self.doJsonRequest("POST", "/v1/series",
		reqPostSeries{Series: series}, nil)
}

// GetMetric represents a collection of datapoints returned for a given query. It
// is subtly different than a post metric so we can't reuse Metric here
type GetMetric struct {
	Metric    string      `json:"metric`
	Pointlist []DataPoint `json:"pointlist"`
	Interval  int64       `json:"interval"`
}

// reqQuery from /api/v1/query
type reqQuery struct {
	Series []GetMetric `json:"series"`
}

func (self *Client) GetMetrics(from, to int64, query string) (*reqQuery, error) {
	var out reqQuery
	uri := fmt.Sprintf("/v1/query?from=%d&to=%d&query=%s", from, to, query)
	err := self.doJsonRequest("GET", uri, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
