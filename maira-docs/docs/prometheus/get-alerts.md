# prometheus get-alerts

Get the list of active alerts

## Description

Returns a list of active alerts

## Synopsis

`prometheus get-alerts [--site  <site>] [--cluster  <cluster>]`

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
! prometheus get-alerts --site "localhost"
```
Output: 
```
{
  "status": "success",
  "data": {
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
    ]
  }
}
```

