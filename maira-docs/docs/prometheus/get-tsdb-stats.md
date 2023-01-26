# prometheus get-tsdb-stats

Get timeseries database statistics

## Description

Returns various cardinality statistics about the Prometheus TSDB

## Synopsis

`prometheus get-tsdb-stats [--site  <site>] [--cluster  <cluster>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of prometheus cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "prometheus-default"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `prometheus-default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! prometheus get-tsdb-stats --site "localhost" 
```
Output: 
```
{
  "status": "success",
  "data": {
    "headStats": {
      "numSeries": 915,
      "numLabelPairs": 368,
      "chunkCount": 2457,
      "minTime": 1634810400284,
      "maxTime": 1634814150284
    },
    "seriesCountByMetricName": [
      {
        "name": "prometheus_http_request_duration_seconds_bucket",
        "value": 210
      },
      {
        "name": "prometheus_http_response_size_bytes_bucket",
        "value": 189
      },
      {
        "name": "prometheus_http_requests_total",
        "value": 23
      },
      {
        "name": "prometheus_http_response_size_bytes_count",
        "value": 21
      },
      {
        "name": "prometheus_http_request_duration_seconds_sum",
        "value": 21
      },
      {
        "name": "prometheus_http_response_size_bytes_sum",
        "value": 21
      },
      {
        "name": "prometheus_http_request_duration_seconds_count",
        "value": 21
      },
      {
        "name": "net_conntrack_dialer_conn_failed_total",
        "value": 20
      },
      {
        "name": "prometheus_sd_kubernetes_events_total",
        "value": 18
      },
      {
        "name": "prometheus_tsdb_compaction_duration_seconds_bucket",
        "value": 15
      }
    ],
    "labelValueCountByLabelName": [
      {
        "name": "__name__",
        "value": 210
      },
      {
        "name": "le",
        "value": 70
      },
      {
        "name": "handler",
        "value": 21
      },
      {
        "name": "quantile",
        "value": 9
      },
      {
        "name": "role",
        "value": 6
      },
      {
        "name": "dialer_name",
        "value": 5
      },
      {
        "name": "code",
        "value": 5
      },
      {
        "name": "slice",
        "value": 4
      },
      {
        "name": "config",
        "value": 4
      },
      {
        "name": "reason",
        "value": 4
      }
    ],
    "memoryInBytesByLabelName": [
      {
        "name": "__name__",
        "value": 7966
      },
      {
        "name": "handler",
        "value": 337
      },
      {
        "name": "le",
        "value": 323
      },
      {
        "name": "rule_group",
        "value": 117
      },
      {
        "name": "dialer_name",
        "value": 47
      },
      {
        "name": "slice",
        "value": 43
      },
      {
        "name": "role",
        "value": 43
      },
      {
        "name": "instance",
        "value": 42
      },
      {
        "name": "revision",
        "value": 40
      },
      {
        "name": "config",
        "value": 36
      }
    ],
    "seriesCountByLabelValuePair": [
      {
        "name": "instance=localhost:9090",
        "value": 901
      },
      {
        "name": "job=prometheus",
        "value": 901
      },
      {
        "name": "__name__=prometheus_http_request_duration_seconds_bucket",
        "value": 210
      },
      {
        "name": "__name__=prometheus_http_response_size_bytes_bucket",
        "value": 189
      },
      {
        "name": "le=+Inf",
        "value": 47
      },
      {
        "name": "handler=/api/v1/series",
        "value": 25
      },
      {
        "name": "handler=/api/v1/query_range",
        "value": 25
      },
      {
        "name": "handler=/manifest.json",
        "value": 24
      },
      {
        "name": "handler=/api/v1/query",
        "value": 24
      },
      {
        "name": "handler=/api/v1/labels",
        "value": 24
      }
    ]
  }
}
```

