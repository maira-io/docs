# prometheus get-targets

Get the Prometheus target discovery

## Description

Returns an overview of the current state of the Prometheus target discovery

## Synopsis

`prometheus get-targets [--site  <site>] [--cluster  <cluster>] [<state>]`

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


#### `state` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Used to filter targets (either active or dropped). Possible values are active, dropped, any  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"active"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! prometheus get-targets --site "localhost" "active"
```
Output: 
```
{
  "status": "success",
  "data": {
    "activeTargets": [
      {
        "discoveredLabels": {
          "__address__": "localhost:1234",
          "__metrics_path__": "/metrics",
          "__scheme__": "http",
          "job": "kafka"
        },
        "labels": {
          "instance": "localhost:1234",
          "job": "kafka"
        },
        "scrapePool": "kafka",
        "scrapeUrl": "http://localhost:1234/metrics",
        "globalUrl": "http://beehyv-H81M-S:1234/metrics",
        "lastError": "Get \"http://localhost:1234/metrics\": dial tcp 127.0.0.1:1234: connect: connection refused",
        "lastScrape": "2021-10-21T08:02:32.877553278+05:30",
        "lastScrapeDuration": 0.000657934,
        "health": "down"
      },
      {
        "discoveredLabels": {
          "__address__": "localhost:9100",
          "__metrics_path__": "/metrics",
          "__scheme__": "http",
          "job": "node-exporter"
        },
        "labels": {
          "instance": "localhost:9100",
          "job": "node-exporter"
        },
        "scrapePool": "node-exporter",
        "scrapeUrl": "http://localhost:9100/metrics",
        "globalUrl": "http://beehyv-H81M-S:9100/metrics",
        "lastError": "Get \"http://localhost:9100/metrics\": dial tcp 127.0.0.1:9100: connect: connection refused",
        "lastScrape": "2021-10-21T08:02:30.284777044+05:30",
        "lastScrapeDuration": 0.000598593,
        "health": "down"
      },
      {
        "discoveredLabels": {
          "__address__": "localhost:9090",
          "__metrics_path__": "/metrics",
          "__scheme__": "http",
          "job": "prometheus"
        },
        "labels": {
          "instance": "localhost:9090",
          "job": "prometheus"
        },
        "scrapePool": "prometheus",
        "scrapeUrl": "http://localhost:9090/metrics",
        "globalUrl": "http://beehyv-H81M-S:9090/metrics",
        "lastError": "",
        "lastScrape": "2021-10-21T08:02:39.829747752+05:30",
        "lastScrapeDuration": 0.009807423,
        "health": "up"
      }
    ],
    "droppedTargets": []
  }
}
```

