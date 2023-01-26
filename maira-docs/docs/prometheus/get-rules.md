# prometheus get-rules

Get alerting and recording rule

## Description

Returns a list of alerting and recording rules that are currently loaded. In addition it returns the currently active alerts fired by the Prometheus instance of each alerting rule.

## Synopsis

`prometheus get-rules [--site  <site>] [--cluster  <cluster>] [--type  <type>]`

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


#### `type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Used to filter rules (either alert or record). Possible values are alert, record  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--type  "alert"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! prometheus get-rules --site "localhost" --type "alert"
```
Output: 
```
{
  "status": "success",
  "data": {
    "groups": [
      {
        "name": "alert_rules",
        "file": "/usr/local/bin/prometheus/prometheus_rules.yml",
        "rules": [
          {
            "state": "firing",
            "name": "InstanceDown",
            "query": "up == 0",
            "duration": 15,
            "labels": {
              "severity": "critical"
            },
            "annotations": {
              "description": "[{{ $labels.instance }}] of job [{{ $labels.job }}] has been down for more than 1 minute.",
              "summary": "Instance [{{ $labels.instance }}] down"
            },
            "alerts": [
              {
                "labels": {
                  "alertname": "InstanceDown",
                  "instance": "localhost:9100",
                  "job": "node-exporter",
                  "severity": "critical"
                },
                "annotations": {
                  "description": "[localhost:9100] of job [node-exporter] has been down for more than 1 minute.",
                  "summary": "Instance [localhost:9100] down"
                },
                "state": "firing",
                "activeAt": "2021-10-13T12:05:20.566139241Z",
                "value": "0e+00"
              },
              {
                "labels": {
                  "alertname": "InstanceDown",
                  "instance": "localhost:1234",
                  "job": "kafka",
                  "severity": "critical"
                },
                "annotations": {
                  "description": "[localhost:1234] of job [kafka] has been down for more than 1 minute.",
                  "summary": "Instance [localhost:1234] down"
                },
                "state": "firing",
                "activeAt": "2021-10-13T12:05:20.566139241Z",
                "value": "0e+00"
              }
            ],
            "health": "ok",
            "evaluationTime": 0.001672296,
            "lastEvaluation": "2021-10-21T08:07:35.567271977+05:30",
            "type": "alerting"
          }
        ],
        "interval": 15,
        "evaluationTime": 0.001713941,
        "lastEvaluation": "2021-10-21T08:07:35.567239362+05:30"
      },
      {
        "name": "custom_rules",
        "file": "/usr/local/bin/prometheus/prometheus_rules.yml",
        "rules": [],
        "interval": 15,
        "evaluationTime": 0.000517493,
        "lastEvaluation": "2021-10-21T08:07:44.069116798+05:30"
      }
    ]
  }
}
```

